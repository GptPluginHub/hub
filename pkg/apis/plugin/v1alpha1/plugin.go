package v1alpha1

import (
	"context"

	"github.com/GptPluginHub/hub/pkg/application"
	"github.com/GptPluginHub/hub/pkg/config"

	"google.golang.org/protobuf/types/known/emptypb"
	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
)

type PluginHandler struct {
	*pluginv1alpha1.UnimplementedPluginServiceServer
	PluginAppInterface      application.PluginAppInterface
	PluginScoreAppInterface application.PluginScoreAppInterface
}

func (p *PluginHandler) CreatePlugin(ctx context.Context, request *pluginv1alpha1.CreatePluginRequest) (*emptypb.Empty, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, err
	}
	err := p.PluginAppInterface.CreatePlugin(ctx, request)
	return &emptypb.Empty{}, err
}

func (p *PluginHandler) ListPlugins(ctx context.Context, req *pluginv1alpha1.ListPluginRequest) (*pluginv1alpha1.ListPluginResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	plugins, err := p.PluginAppInterface.ListPlugins(ctx, req)
	return plugins, err
}

func (p *PluginHandler) CreatePluginScore(ctx context.Context, req *pluginv1alpha1.CreatePluginScoreRequest) (*emptypb.Empty, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, err
	}
	err := p.PluginScoreAppInterface.CreatePluginScore(ctx, req)
	return &emptypb.Empty{}, err
}

var _ pluginv1alpha1.PluginServiceServer = new(PluginHandler)

func NewPluginHandler(cfg config.Config) pluginv1alpha1.PluginServiceServer {
	app := application.NewPluginApp(cfg)
	pluginScoreApp := application.NewPluginScoreAppInterface(cfg)
	return &PluginHandler{PluginAppInterface: app, PluginScoreAppInterface: pluginScoreApp}
}
