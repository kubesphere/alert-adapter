package monitoring

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"kubesphere.io/alert/pkg/logger"
)

var client = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   10000,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

const (
	DefaultScheme = "http"
)

func SendMonitoringRequest(uriPath string, extraQueryParams string, metrics []string) string {
	var params = url.Values{
		"metrics_filter": []string{strings.Join(metrics, "|")},
	}

	var urlStr string

	queryParams := params.Encode()
	if extraQueryParams != "" {
		queryParams = queryParams + "&" + extraQueryParams
	}

	if strings.Contains(uriPath, "?") {
		if strings.HasPrefix(uriPath, "http://") || strings.HasPrefix(uriPath, "https://") {
			urlStr = uriPath + "&" + queryParams
		} else {
			urlStr = DefaultScheme + "://" + uriPath + "&" + queryParams
		}
	} else {
		if strings.HasPrefix(uriPath, "http://") || strings.HasPrefix(uriPath, "https://") {
			urlStr = uriPath + "?" + queryParams
		} else {
			urlStr = DefaultScheme + "://" + uriPath + "?" + queryParams
		}
	}

	logger.Info(nil, "SendMonitoringRequest %s", urlStr)

	response, err := client.Get(urlStr)
	if err != nil {
		logger.Error(nil, "SendMonitoringRequest get error:", err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			logger.Error(nil, "SendMonitoringRequest read error:", err.Error())
		}

		return string(contents)
	}
	return ""
}
