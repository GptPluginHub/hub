package application

import (
	"context"
	"errors"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/domain"

	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
	"hub.io/api/types"
	"k8s.io/klog"
)

type PluginAppInterface interface {
	ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error)
	CreatePlugin(ctx context.Context, req *pluginv1alpha1.CreatePluginRequest) error
}

var _ PluginAppInterface = new(PluginApp)

type PluginApp struct {
	domain.Plugin
	domain.PluginMetadata
}

func (p *PluginApp) CreatePlugin(ctx context.Context, req *pluginv1alpha1.CreatePluginRequest) error {
	klog.Info("CreatePlugin domain: ", req.Domain)
	aiPluginURL := p.GeneratePluginURL(ctx, req.Domain)
	klog.Info("CreatePlugin aiPluginURL: ", aiPluginURL)
	aiPluginInfo, err := p.FetchAiPluginInfo(aiPluginURL)
	if err != nil {
		return err
	}
	pluginModel := p.GeneratePluginModel(ctx, req.Domain, req.Label, aiPluginInfo)
	klog.Info("CreatePlugin pluginModel: ", pluginModel)
	if p.CheckExist(ctx, pluginModel.Name) {
		return errors.New("plugin already exist")
	}
	if err = p.AddPlugin(ctx, &pluginModel); err != nil {
		return err
	}
	plugin, err := p.PluginInfra.GetPluginByName(ctx, pluginModel.Name)
	if err != nil {
		klog.Errorf("CreatePlugin GetPluginByName error: %v", err)
		return err
	}
	pluginMetadata := p.GeneratePluginMetadata(ctx, plugin.ID, aiPluginURL)
	if err = p.PluginMetadata.AddPluginMetadata(ctx, pluginMetadata); err != nil {
		klog.Errorf("CreatePlugin AddPluginMetadata error: %v", err)
		return err
	}
	return nil
}

func (p *PluginApp) ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error) {
	page := &types.Page{
		PageSize: req.PageSize,
		Page:     req.Page,
	}
	if page.PageSize == 0 {
		page.PageSize = 10
	}
	if req.SortByFieldName == "" {
		req.SortByFieldName = "created_at"
	}
	plugins, err := p.Plugin.ListPluginByFuzzyName(ctx, req.FuzzyName, req.SortByFieldName, req.OrderBy, page)
	if err != nil {
		klog.Errorf("ListPlugins error: %v", err)
		return nil, err
	}
	pluginList := make([]*pluginv1alpha1.Plugin, 0)
	for _, plugin := range plugins {
		pluginList = append(pluginList, ModelPluginConvToAPIPlugin(*plugin))
	}
	return &pluginv1alpha1.ListPluginResponse{
		Item: pluginList,
		Page: page,
	}, nil
}

func NewPluginApp(cfg config.Config) PluginAppInterface {
	plugin, err := domain.NewPlugin(cfg.MysqlOptions)
	if err != nil {
		panic(err)
	}
	pluginMetadata, err := domain.NewPluginMetadata(cfg.MysqlOptions)
	if err != nil {
		panic(err)
	}
	return &PluginApp{Plugin: *plugin, PluginMetadata: *pluginMetadata}
}
