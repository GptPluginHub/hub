package application

import (
	"github.com/GptPluginHub/hub/pkg/model"

	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
)

func ModelPluginConvToAPIPlugin(plugin model.Plugin) *pluginv1alpha1.Plugin {
	p := pluginv1alpha1.Plugin{
		Id:          int32(plugin.ID),
		Name:        plugin.Name,
		Description: plugin.Description,
		LogoUrl:     plugin.LogoURL,
		ApiUrl:      plugin.APIURL,
		InstallNum:  int32(plugin.InstallNum),
		Score:       float32(plugin.Score),
		Heat:        int32(plugin.Heat),
		CreatedAt:   plugin.CreatedAt,
		UpdatedAt:   plugin.UpdatedAt,
		Labels:      plugin.Label,
	}
	return &p
}
