package kubesphere

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "kubesphere.io/alert/pkg/client/kubernetes"
	"kubesphere.io/alert/pkg/client/monitoring"
	"kubesphere.io/alert/pkg/config"
	"kubesphere.io/alert/pkg/logger"
	"kubesphere.io/alert/pkg/metric"
)

// Resource
type Resource struct {
	MetricURI     string
	RsTypeName    string
	Namespace     string
	Node          string
	WorkloadKind  string
	LabelSelector string
}

func getKV(kv map[string]string) (string, string) {
	key := ""
	value := ""
	for k, v := range kv {
		key = k
		value = v
	}

	return key, value
}

func assembleMetricURL(monitoringHost string, uriTmpls []string, rsFilterURI map[string]string) (string, error) {
	if len(rsFilterURI) == 0 {
		for _, uriTmpl := range uriTmpls {
			if !strings.Contains(uriTmpl, "{") {
				uri := fmt.Sprintf("%s%s", monitoringHost, uriTmpl)
				return uri, nil
			}
		}
	} else {
		r, err := regexp.Compile(`\{\w+\}`)
		if err != nil {
			return "", err
		}

		for _, uriTmpl := range uriTmpls {
			if len(rsFilterURI) != strings.Count(uriTmpl, "{") {
				continue
			}

			uri := r.ReplaceAllStringFunc(uriTmpl, func(s string) string {
				s = strings.Trim(s, "{")
				s = strings.Trim(s, "}")
				return rsFilterURI[s]
			})

			if !strings.Contains(uri, "{") {
				uri = fmt.Sprintf("%s%s", monitoringHost, uri)
				return uri, nil
			}
		}
	}

	return "", nil
}

func assembleSelectorURL(monitoringHost string, uriTmpls []string, rsFilterURI map[string]string) (string, error) {
	r, err := regexp.Compile(`\{\w+\}`)
	if err != nil {
		return "", err
	}
	for _, uriTmpl := range uriTmpls {
		index := strings.Index(uriTmpl, "?")
		if index >= 0 {
			uriTmpl = uriTmpl[0:index]
		}
		uri := r.ReplaceAllStringFunc(uriTmpl, func(s string) string {
			s = strings.Trim(s, "{")
			s = strings.Trim(s, "}")
			return rsFilterURI[s]
		})
		if !strings.Contains(uri, "{") {
			uri = fmt.Sprintf("%s%s?resources_filter=", monitoringHost, uri)
			return uri, nil
		}
	}

	return "", nil
}

func parseResource(rsTypeName string, rsTypeParam string, rsFilterName string, rsFilterParam string) *Resource {
	resource := &Resource{}

	resource.MetricURI = ""
	resource.Namespace = ""
	resource.LabelSelector = ""

	resource.RsTypeName = rsTypeName

	uriTmpls := strings.Split(rsTypeParam, ",")

	rsFilterURI := make(map[string]string)
	err := json.Unmarshal([]byte(rsFilterParam), &rsFilterURI)
	if err != nil {
		logger.Error(nil, "Unmarshal rsFilterURI Error: %v", err)
	}

	cfg := config.GetInstance()
	monitoringHost := cfg.App.MonitoringHost
	delete(rsFilterURI, "monitoring_host")

	if rsFilterURI["selector"] == "" {
		resource.MetricURI, err = assembleMetricURL(monitoringHost, uriTmpls, rsFilterURI)
		if err != nil {
			logger.Error(nil, "assembleMetricURL Error: %v", err)
		}
	} else {
		resourceSelector := []map[string]string{}
		err := json.Unmarshal([]byte(rsFilterURI["selector"]), &resourceSelector)
		if err != nil {
			logger.Error(nil, "Unmarshal resourceSelector Error: %v", err)
		}

		resource.MetricURI, err = assembleSelectorURL(monitoringHost, uriTmpls, rsFilterURI)
		if err != nil {
			logger.Error(nil, "assembleSelectorURL Error: %v", err)
		}
		for _, kv := range resourceSelector {
			k, v := getKV(kv)
			resource.LabelSelector = resource.LabelSelector + fmt.Sprintf("%s=%s,", k, v)
		}
		resource.LabelSelector = strings.TrimSuffix(resource.LabelSelector, ",")

		switch resource.RsTypeName {
		case "workload":
			resource.Namespace = rsFilterURI["ns_name"]
			resource.WorkloadKind = rsFilterURI["workload_kind"]
		case "pod":
			resource.Namespace = rsFilterURI["ns_name"]
			resource.Node = rsFilterURI["node_id"]
		}
	}

	return resource
}

func getResourceFilterURIBySelector(resource *Resource) (string, bool) {
	//TODO: Get resources by label selector from all resource types.
	switch resource.RsTypeName {
	case "node":
		nodeList, err := k8sclient.NewK8sClient().CoreV1().Nodes().List(metaV1.ListOptions{LabelSelector: resource.LabelSelector})
		if err != nil {
			logger.Error(nil, "getResourceFilterURIBySelector list nodes error: %v", err)
			return resource.MetricURI, false
		}

		if len(nodeList.Items) == 0 {
			return "", false
		}

		resources := ""
		for _, node := range nodeList.Items {
			resources = resources + node.Name + "|"
		}
		resources = strings.TrimSuffix(resources, "|")
		return resource.MetricURI + resources, true
	case "workload":
		switch resource.WorkloadKind {
		case "deployment":
			deploymentList, err := k8sclient.NewK8sClient().ExtensionsV1beta1().Deployments(resource.Namespace).List(metaV1.ListOptions{LabelSelector: resource.LabelSelector})
			if err != nil {
				logger.Error(nil, "getResourceFilterURIBySelector list deployments error: %v", err)
				return resource.MetricURI, false
			}

			if len(deploymentList.Items) == 0 {
				return "", false
			}

			resources := ""
			for _, deployment := range deploymentList.Items {
				resources = resources + deployment.Name + "|"
			}
			resources = strings.TrimSuffix(resources, "|")
			return resource.MetricURI + resources, true
		case "statefulset":
			statefulsetList, err := k8sclient.NewK8sClient().AppsV1().StatefulSets(resource.Namespace).List(metaV1.ListOptions{LabelSelector: resource.LabelSelector})
			if err != nil {
				logger.Error(nil, "getResourceFilterURIBySelector list statefulsets error: %v", err)
				return resource.MetricURI, false
			}

			if len(statefulsetList.Items) == 0 {
				return "", false
			}

			resources := ""
			for _, statefulset := range statefulsetList.Items {
				resources = resources + statefulset.Name + "|"
			}
			resources = strings.TrimSuffix(resources, "|")
			return resource.MetricURI + resources, true
		case "daemonset":
			daemonsetList, err := k8sclient.NewK8sClient().ExtensionsV1beta1().DaemonSets(resource.Namespace).List(metaV1.ListOptions{LabelSelector: resource.LabelSelector})
			if err != nil {
				logger.Error(nil, "getResourceFilterURIBySelector list daemonsets error: %v", err)
				return resource.MetricURI, false
			}

			if len(daemonsetList.Items) == 0 {
				return "", false
			}

			resources := ""
			for _, daemonset := range daemonsetList.Items {
				resources = resources + daemonset.Name + "|"
			}
			resources = strings.TrimSuffix(resources, "|")
			return resource.MetricURI + resources, true
		}
	case "pod":
		if resource.Namespace != "" {
			podList, err := k8sclient.NewK8sClient().CoreV1().Pods(resource.Namespace).List(metaV1.ListOptions{LabelSelector: resource.LabelSelector})
			if err != nil {
				logger.Error(nil, "getResourceFilterURIBySelector list pods error: %v", err)
				return resource.MetricURI, false
			}

			if len(podList.Items) == 0 {
				return "", false
			}

			resources := ""
			for _, pod := range podList.Items {
				resources = resources + pod.Name + "|"
			}
			resources = strings.TrimSuffix(resources, "|")
			return resource.MetricURI + resources, true
		}
		if resource.Node != "" {
			podList, err := k8sclient.NewK8sClient().CoreV1().Pods("").List(metaV1.ListOptions{FieldSelector: "spec.nodeName=" + resource.Node, LabelSelector: resource.LabelSelector})
			if err != nil {
				logger.Error(nil, "getResourceFilterURIBySelector list pods error: %v", err)
				return resource.MetricURI, false
			}

			if len(podList.Items) == 0 {
				return "", false
			}

			resources := ""
			for _, pod := range podList.Items {
				resources = resources + pod.Name + "|"
			}
			resources = strings.TrimSuffix(resources, "|")
			return resource.MetricURI + resources, true
		}
	default:
		return resource.MetricURI, false
	}

	return resource.MetricURI, false
}

// Metric
const (
	MetricTypeVector = "vector"
	MetricTypeMatrix = "matrix"
)

type Metrics struct {
	MetricsLevel string   `json:"metrics_level"`
	Results      []Metric `json:"results"`
}

type Metric struct {
	MetricName string     `json:"metric_name, omitempty"`
	Status     string     `json:"status"`
	Data       MetricData `json:"data, omitempty"`
}

type MetricData struct {
	Result     json.RawMessage
	ResultType string `json:"resultType"`
}

type InstantMetric struct {
	MetricInfo  map[string]interface{} `json:"metric"`
	MetricValue []interface{}          `json:"value"`
}

type TimeSeriesMetric struct {
	MetricInfo   map[string]interface{} `json:"metric"`
	MetricValues [][]interface{}        `json:"values"`
}

func getResourceTimeSeriesMetric(metricToRule map[string][]string, metricName string, metricStr string) []metric.ResourceMetrics {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(nil, "getResourceTimeSeriesMetric panic:", err)
		}
	}()

	var metrics Metrics
	err := json.Unmarshal([]byte(metricStr), &metrics)
	if err != nil {
		logger.Error(nil, "getResourceTimeSeriesMetric unmarshal metric string error:", err.Error())
	}
	l := len(metrics.Results)

	rms := []metric.ResourceMetrics{}

	for i := 0; i < l; i++ {
		if metrics.Results[i].MetricName != metricName {
			continue
		}

		tp := metrics.Results[i].Data.ResultType
		r := metrics.Results[i].Data.Result
		if tp == MetricTypeVector {
			var instantMetrics []InstantMetric

			d := json.NewDecoder(bytes.NewReader(r))
			d.UseNumber()
			if err := d.Decode(&instantMetrics); err != nil {
				logger.Error(nil, "getResourceTimeSeriesMetric decode metric vector error:", err.Error())
			}

			item := make(map[string][]metric.TV)

			for i := 0; i < len(instantMetrics); i++ {
				metricInfo := instantMetrics[i].MetricInfo
				resourceName := metricInfo["resource_name"].(string)

				metricValue := instantMetrics[i].MetricValue
				t, err := metricValue[0].(json.Number).Float64()
				if err != nil {
					logger.Error(nil, "getResourceTimeSeriesMetric read metric timestamp error:", err)
					continue
				}
				v := metricValue[1].(string)
				tv := metric.TV{T: int64(t), V: v}
				item[resourceName] = []metric.TV{tv}
			}

			for _, ruleId := range metricToRule[metricName] {
				rm := metric.ResourceMetrics{
					RuleId:         ruleId,
					MetricName:     metricName,
					ResourceMetric: item,
				}

				rms = append(rms, rm)
			}
		} else if tp == MetricTypeMatrix {
			var timeRangeMetrics []TimeSeriesMetric

			d := json.NewDecoder(bytes.NewReader(r))
			d.UseNumber()
			if err := d.Decode(&timeRangeMetrics); err != nil {
				logger.Error(nil, "getResourceTimeSeriesMetric decode metric matrix error:", err.Error())
			}

			item := make(map[string][]metric.TV)

			for i := 0; i < len(timeRangeMetrics); i++ {
				metricItem := timeRangeMetrics[i]
				metricInfo := metricItem.MetricInfo
				resourceName := metricInfo["resource_name"].(string)
				metricValues := metricItem.MetricValues
				var tvs []metric.TV
				for i := 0; i < len(metricValues); i++ {
					t, err := metricValues[i][0].(json.Number).Int64()
					if err != nil {
						logger.Error(nil, "getResourceTimeSeriesMetric read metric timestamp error:", err)
						continue
					}
					v := metricValues[i][1].(string)
					tv := metric.TV{T: t, V: v}
					tvs = append(tvs, tv)
				}
				if len(tvs) != 0 {
					item[resourceName] = tvs
				}
			}

			for _, ruleId := range metricToRule[metricName] {
				rm := metric.ResourceMetrics{
					RuleId:         ruleId,
					MetricName:     metricName,
					ResourceMetric: item,
				}

				rms = append(rms, rm)
			}
		}
	}
	return rms
}

func GetMetric(rsTypeName string, rsTypeParam string, rsFilterName string, rsFilterParam string, extraQueryParams string, metrics []string, metricToRule map[string][]string) string {
	resource := parseResource(rsTypeName, rsTypeParam, rsFilterName, rsFilterParam)

	metricStr := ""
	if resource.LabelSelector == "" {
		metricStr = monitoring.SendMonitoringRequest(resource.MetricURI, extraQueryParams, metrics)
	} else {
		MetricURI, hasResource := getResourceFilterURIBySelector(resource)
		if !hasResource {
			return ""
		}
		metricStr = monitoring.SendMonitoringRequest(MetricURI, extraQueryParams, metrics)
	}

	resourceMetricsAll := []metric.ResourceMetrics{}

	for _, mt := range metrics {
		resourceMetrics := getResourceTimeSeriesMetric(metricToRule, mt, metricStr)
		for _, rm := range resourceMetrics {
			resourceMetricsAll = append(resourceMetricsAll, rm)
		}
	}

	resourceMetricsAllStr, _ := json.Marshal(resourceMetricsAll)

	return string(resourceMetricsAllStr)
}
