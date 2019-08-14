package kubesphere

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/emicklei/go-restful"

	"kubesphere.io/alert/pkg/constants"
	"kubesphere.io/alert/pkg/logger"
	"kubesphere.io/alert/pkg/metric"
	"kubesphere.io/alert/pkg/notification"
)

func parseBool(input string) bool {
	if input == "true" {
		return true
	} else {
		return false
	}
}

func GetMetrics(request *restful.Request, response *restful.Response) {
	metricParamStr := request.QueryParameter("metric_param")

	metricParam := metric.MetricParam{}

	err := json.Unmarshal([]byte(metricParamStr), &metricParam)
	if err != nil {
		logger.Error(nil, "Unmarshal Metric Param error: %v", err)
		return
	}

	resp := GetMetric(metricParam.RsTypeName, metricParam.RsTypeParam, metricParam.RsFilterName, metricParam.RsFilterParam, metricParam.ExtraQueryParams, metricParam.Metrics, metricParam.MetricToRule)

	response.Write([]byte(resp))
}

func GetEmail(request *restful.Request, response *restful.Response) {
	resume := parseBool(request.QueryParameter("resume"))
	notificationParamStr := request.QueryParameter("notification_param")

	notificationParam := notification.NotificationParam{}

	err := json.Unmarshal([]byte(notificationParamStr), &notificationParam)
	if err != nil {
		logger.Error(nil, "Unmarshal Notification Param error: %v", err)
		return
	}

	tmpl := template.New("email")
	if resume {
		tmpl.Parse(constants.EmailKubeSphereNotifyResumeTemplate)
	} else {
		tmpl.Parse(constants.EmailKubeSphereNotifyActiveTemplate)
	}

	var content bytes.Buffer
	tmpl.Execute(&content, notificationParam)

	emailContent := notification.EmailContent{
		Html: content.String(),
	}

	emailContentBytes, err := json.Marshal(emailContent)

	if err != nil {
		logger.Error(nil, "Marshal Notification Content error: %v", err)
		return
	}

	email := notification.Email{
		Title:   "KubeSphere Notification",
		Content: string(emailContentBytes),
	}

	response.WriteAsJson(email)
}
