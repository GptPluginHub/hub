package db

import (
	"context"
	"database/sql"

	"github.com/GptPluginHub/hub/pkg/model"

	"k8s.io/klog"
)

type PluginInfraInterface interface {
	// GetPluginByID returns plugin by id
	GetPluginByID(ctx context.Context, id string) (*model.Plugin, error)
	// GetPluginByDomain returns plugin by name
	GetPluginByDomain(ctx context.Context, name string) (*model.Plugin, error)
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

func (m *pluginInfra) GetPluginByDomain(ctx context.Context, name string) (*model.Plugin, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM plugin WHERE domain = ?", name)
	var plugin model.Plugin
	err := row.Scan(&plugin.ID, &plugin.Domain, &plugin.Name, &plugin.Description, &plugin.AuthType, &plugin.LogoURL, &plugin.ContactEmail, &plugin.Organization, &plugin.APIType, &plugin.APIURL, &plugin.Label, &plugin.State, &plugin.InstallNum, &plugin.Score, &plugin.Heat, &plugin.CreatedAt, &plugin.UpdatedAt)
	return &plugin, err
}

func (m *pluginInfra) ListPlugins(ctx context.Context, page model.Page, sortBy model.SortBy, orderBy model.OrderBy, fuzzyName ...string) ([]*model.Plugin, error) {
	panic("implement me")
}

func (m *pluginInfra) CreatePlugin(ctx context.Context, plugin *model.Plugin) error {
	r, err := m.db.ExecContext(ctx, "INSERT INTO plugin (id, domain, name, description, auth_type, logo_url, contact_email,organization, api_type,api_url, label, state, install_num,score, heat, created_at, updated_at) VALUEs (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", plugin.ID, plugin.Domain, plugin.Name, plugin.Description, plugin.AuthType, plugin.LogoURL, plugin.ContactEmail, plugin.Organization, plugin.APIType, plugin.APIURL, plugin.Label, plugin.State, plugin.InstallNum, plugin.Score, plugin.Heat, plugin.CreatedAt, plugin.UpdatedAt)
	if err != nil {
		klog.Errorf("create plugin failed, insert err: %v", err)
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		klog.Errorf("create plugin failed, last_insert_id err: %v", err)
		return err
	}
	klog.Infof("create plugin succ, id: %d", id)
	return nil
}
