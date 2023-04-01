package db

import (
	"context"
	"database/sql"

	"github.com/GptPluginHub/hub/pkg/model"
)

type PluginInfraInterface interface {
	// GetPluginByID returns plugin by id
	GetPluginByID(ctx context.Context, id string) (*model.Plugin, error)
	// GetPluginByName returns plugin by name
	GetPluginByName(ctx context.Context, name string) (*model.Plugin, error)
	// ListPlugins returns all plugins
	ListPlugins(ctx context.Context, page model.Page, sortBy model.SortBy, orderBy model.OrderBy, fuzzyName ...string) ([]*model.Plugin, error)
	// CreatePlugin creates a plugin
	CreatePlugin(ctx context.Context, plugin *model.Plugin) error
}

var _ PluginInfraInterface = new(pluginInfra)

type pluginInfra struct {
	db *sql.DB
}

func NewPluginInfra(db *sql.DB) PluginInfraInterface {
	return &pluginInfra{db}
}

func (m *pluginInfra) GetPluginByID(ctx context.Context, id string) (*model.Plugin, error) {
	panic("implement me")
}

func (m *pluginInfra) GetPluginByName(ctx context.Context, name string) (*model.Plugin, error) {
	panic("implement me")
}

func (m *pluginInfra) ListPlugins(ctx context.Context, page model.Page, sortBy model.SortBy, orderBy model.OrderBy, fuzzyName ...string) ([]*model.Plugin, error) {
	panic("implement me")
}

func (m *pluginInfra) CreatePlugin(ctx context.Context, plugin *model.Plugin) error {
	panic("implement me")
}
