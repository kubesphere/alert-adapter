package kubesphere

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"kubesphere.io/alert/pkg/config"
	"kubesphere.io/alert/pkg/constants"
)

func WebService() *restful.WebService {
	restful.RegisterEntityAccessor(constants.MIME_MERGEPATCH, restful.NewEntityAccessorJSON(restful.MIME_JSON))

	ws := new(restful.WebService)
	ws.Path("/api/v1").Consumes(restful.MIME_JSON, constants.MIME_MERGEPATCH).Produces(restful.MIME_JSON)

	tags := []string{"Metric apis"}

	ws.Route(ws.GET("/metric").To(GetMetrics).
		Doc("Get Metric Values").
		Param(ws.QueryParameter("metric_param", "metric param specify").DataType("string").Required(false)).
		Metadata(restfulspec.KeyOpenAPITags, tags)).
		Consumes(restful.MIME_JSON, constants.MIME_MERGEPATCH).
		Produces(restful.MIME_JSON)

	tags = []string{"Email apis"}

	ws.Route(ws.GET("/email").To(GetEmail).
		Doc("Get Email Text").
		Param(ws.QueryParameter("notification_param", "notification param specify").DataType("string").Required(false)).
		Metadata(restfulspec.KeyOpenAPITags, tags)).
		Consumes(restful.MIME_JSON, constants.MIME_MERGEPATCH).
		Produces(restful.MIME_JSON)

	return ws
}

func Run() {
	cfg := config.GetInstance().LoadConf()

	restful.DefaultContainer.Add(WebService())
	enableCORS()

	listen := fmt.Sprintf(":%s", cfg.App.AdapterPort)

	log.Print(http.ListenAndServe(listen, nil))
}

func enableCORS() {
	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		AllowedDomains: []string{"*"},
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)
}
