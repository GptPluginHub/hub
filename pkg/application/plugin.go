package application

import (
	"context"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/domain"

	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
)

type PluginAppInterface interface {
	ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error)
}

var _ PluginAppInterface = new(PluginApp)

type PluginApp struct {
	domain.Plugin
}

func (p *PluginApp) ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error) {
	panic("implement me")
}

func NewPluginApp(cfg config.Config) PluginAppInterface {
	plugin, err := domain.NewPlugin(cfg.MysqlOptions)
	if err != nil {
		panic(err)
	}
	return &PluginApp{Plugin: *plugin}
}
