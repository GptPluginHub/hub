package application

import (
	"context"
	"errors"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/domain"

	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
	"k8s.io/klog"
)

type PluginAppInterface interface {
	ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error)
	CreatePlugin(ctx context.Context, req *pluginv1alpha1.CreatePluginRequest) error
}

var _ PluginAppInterface = new(PluginApp)

type PluginApp struct {
	domain.Plugin
}

func (p *PluginApp) CreatePlugin(ctx context.Context, req *pluginv1alpha1.CreatePluginRequest) error {
	if p.CheckExist(ctx, req.Domain) {
		return errors.New("plugin already exist")
	}
	klog.Info("CreatePlugin domain: ", req.Domain)
	aiPluginURL := p.GeneratePluginURL(ctx, req.Domain)
	klog.Info("CreatePlugin aiPluginURL: ", aiPluginURL)
	aiPluginInfo, err := p.FetchAiPluginInfo(aiPluginURL)
	if err != nil {
		return err
	}
	pluginModel := p.GeneratePluginModel(ctx, req.Domain, req.Label, aiPluginInfo)
	klog.Info("CreatePlugin pluginModel: ", pluginModel)
	return p.AddPlugin(ctx, &pluginModel)
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
