package v1alpha1

import (
	"context"

	"github.com/GptPluginHub/hub/pkg/application"
	"github.com/GptPluginHub/hub/pkg/config"

	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
)

type PluginHandler struct {
	*pluginv1alpha1.UnimplementedPluginServiceServer
	PluginAppInterface application.PluginAppInterface
}

func (p *PluginHandler) ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error) {
	panic("implement me")
}

var _ pluginv1alpha1.PluginServiceServer = new(PluginHandler)

func NewPluginHandler(cfg config.Config) pluginv1alpha1.PluginServiceServer {
	app := application.NewPluginApp(cfg)
	return &PluginHandler{PluginAppInterface: app}
}
