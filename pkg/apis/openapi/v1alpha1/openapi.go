package v1alpha1

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/GptPluginHub/hub/pkg/application"
)

const (
	SwaggerUITemplateKey         = "openapi_url"
	SwaggerUITemplateValuePrefix = "/api/hub.io/v1alpha1/openapi/data?openapi_url=%s"
	SwaggerUITemplateName        = "template.html"
)

type OpenAPIHandler struct {
	OpenAPIApp application.OpenAPIAppInterface
}

func NewOpenAPIHandler(dir, fileName string) *OpenAPIHandler {
	return &OpenAPIHandler{OpenAPIApp: application.NewOpenAPIApp(dir, fileName)}
}

func (o *OpenAPIHandler) OpenAPIHandler(w http.ResponseWriter, r *http.Request) {
	openapiURL := r.FormValue(SwaggerUITemplateKey)
	if openapiURL == "" {
		return
	}
	openapiJSONURL := map[string]string{
		SwaggerUITemplateKey: fmt.Sprintf(SwaggerUITemplateValuePrefix, openapiURL),
	}
	t := template.New(SwaggerUITemplateName)
	tpl := template.Must(t.Parse(SwaggerUITemplateHTML))
	tpl.Execute(w, openapiJSONURL)
}

func (o *OpenAPIHandler) OpenAPIHandlerData(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue(SwaggerUITemplateKey)
	if value == "" {
		return
	}
	apiFilePath, err := o.OpenAPIApp.GetOpenAPIFilePath(value)
	if err != nil {
		return
	}
	file, err := os.Open(apiFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
