package v1alpha1

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/GptPluginHub/hub/pkg/application"
	"github.com/GptPluginHub/hub/pkg/config"

	"hub.io/api/types"
	"k8s.io/klog"
)

const (
	SwaggerUITemplateKey         = "openapi_url"
	SwaggerUITemplateValuePrefix = "/api/hub.io/v1alpha1/openapi/data?openapi_url=%s"
	SwaggerUITemplateName        = "template.html"
)

type OpenAPIHandler struct {
	OpenAPIApp application.OpenAPIAppInterface
}

func NewOpenAPIHandler(config *config.Config) *OpenAPIHandler {
	return &OpenAPIHandler{OpenAPIApp: application.NewOpenAPIApp(config.APICacheConf.Dir, config.APICacheConf.FileName, config.MysqlOptions)}
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
	apiData, err := o.OpenAPIApp.GetOpenAPIData(r.Context(), value)
	if err != nil {
		return
	}
	defer apiData.Close()
	io.Copy(w, apiData)
}

func (o *OpenAPIHandler) IninPluginMetadata(w http.ResponseWriter, r *http.Request) {
	app := o.OpenAPIApp.(*application.OpenAPIApp)
	plugins, err := app.Plugin.ListPluginByFuzzyName(r.Context(), "", "id", types.OrderBy_desc, &types.Page{
		Page:     1,
		PageSize: 100,
	})
	if err != nil {
		return
	}
	for _, plugin := range plugins {
		url := app.GeneratePluginURL(r.Context(), plugin.Domain)
		pluginMetadata := app.GeneratePluginMetadata(r.Context(), plugin.ID, url)
		if err = app.PluginMetadata.AddPluginMetadata(r.Context(), pluginMetadata); err != nil {
			klog.Errorf("CreatePlugin AddPluginMetadata error: %v", err)
		}
	}
}
