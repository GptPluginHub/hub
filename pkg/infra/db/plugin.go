package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/model"

	"k8s.io/klog"
)

type PluginInfraInterface interface {
	// GetPluginByID returns plugin by id
	GetPluginByID(ctx context.Context, id string) (*model.Plugin, error)
	// GetPluginByName returns plugin by name
	GetPluginByName(ctx context.Context, name string) (*model.Plugin, error)
	// ListPlugins returns all plugins
	ListPlugins(ctx context.Context, limit, offset int32, orderBy, sortFieldName string, fuzzyName ...string) ([]*model.Plugin, error)
	// CreatePlugin creates a plugin
	CreatePlugin(ctx context.Context, plugin *model.Plugin) error
	// CountPlugins returns the number of plugins
	CountPlugins(ctx context.Context, fuzzyName ...string) (int, error)
	// UpdatePluginScoreAndHeat updates plugin state
	UpdatePluginScoreAndHeat(ctx context.Context, id, heat int32, score float64) error
}

var _ PluginInfraInterface = new(pluginInfra)

type pluginInfra struct {
	db *sql.DB
}

func NewPluginInfra(mysqlConf *config.MysqlOptions) (PluginInfraInterface, error) {
	db := dbInit(mysqlConf)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &pluginInfra{db}, nil
}

func (m *pluginInfra) GetPluginByID(ctx context.Context, id string) (*model.Plugin, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM plugin WHERE id = ?", id)
	var plugin model.Plugin
	var labelsStr string
	err := row.Scan(&plugin.ID, &plugin.Domain, &plugin.Name, &plugin.Description, &plugin.AuthType, &plugin.LogoURL, &plugin.ContactEmail, &plugin.Organization, &plugin.APIType, &plugin.APIURL, labelsStr, &plugin.State, &plugin.InstallNum, &plugin.Score, &plugin.Heat, &plugin.CreatedAt, &plugin.UpdatedAt)
	json.Unmarshal([]byte(labelsStr), &plugin.Label)
	return &plugin, err
}

func (m *pluginInfra) GetPluginByName(ctx context.Context, name string) (*model.Plugin, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM plugin WHERE name = ?", name)
	var plugin model.Plugin
	var labelsStr string
	err := row.Scan(&plugin.ID, &plugin.Domain, &plugin.Name, &plugin.Description, &plugin.AuthType, &plugin.LogoURL, &plugin.ContactEmail, &plugin.Organization, &plugin.APIType, &plugin.APIURL, labelsStr, &plugin.State, &plugin.InstallNum, &plugin.Score, &plugin.Heat, &plugin.CreatedAt, &plugin.UpdatedAt)
	json.Unmarshal([]byte(labelsStr), &plugin.Label)
	return &plugin, err
}

func (m *pluginInfra) ListPlugins(ctx context.Context, limit, offset int32, orderBy, sortFieldName string, fuzzyName ...string) ([]*model.Plugin, error) {
	var plugins []*model.Plugin
	var rows *sql.Rows
	var err error
	var sql string
	if len(fuzzyName) > 0 && fuzzyName[0] != "" {
		sql = fmt.Sprintf("SELECT * FROM plugin WHERE name LIKE '%%%s%%' ORDER BY %s %s LIMIT ? OFFSET ?", fuzzyName[0], sortFieldName, orderBy)
	} else {
		sql = fmt.Sprintf("SELECT * FROM plugin ORDER BY %s %s LIMIT ? OFFSET ?", sortFieldName, orderBy)
	}
	rows, err = m.db.QueryContext(ctx, sql, limit, offset)
	if err != nil {
		klog.Errorf("list plugins failed, query err: %v", err)
		return nil, err
	}
	for rows.Next() {
		var plugin model.Plugin
		var labelsStr string
		err := rows.Scan(&plugin.ID, &plugin.Domain, &plugin.Name, &plugin.Description, &plugin.AuthType, &plugin.LogoURL, &plugin.ContactEmail, &plugin.Organization, &plugin.APIType, &plugin.APIURL, &labelsStr, &plugin.State, &plugin.InstallNum, &plugin.Score, &plugin.Heat, &plugin.CreatedAt, &plugin.UpdatedAt)
		if err != nil {
			klog.Errorf("list plugins failed, scan err: %v", err)
			return nil, err
		}
		json.Unmarshal([]byte(labelsStr), &plugin.Label)
		plugins = append(plugins, &plugin)
	}
	return plugins, nil
}

func (m *pluginInfra) CreatePlugin(ctx context.Context, plugin *model.Plugin) error {
	label := "[]"
	if len(plugin.Label) != 0 {
		labelBytes, _ := json.Marshal(plugin.Label)
		label = string(labelBytes)
	}
	r, err := m.db.ExecContext(ctx, "INSERT INTO plugin (id, domain, name, description, auth_type, logo_url, contact_email,organization, api_type,api_url, label, state, install_num,score, heat, created_at, updated_at) VALUEs (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", plugin.ID, plugin.Domain, plugin.Name, plugin.Description, plugin.AuthType, plugin.LogoURL, plugin.ContactEmail, plugin.Organization, plugin.APIType, plugin.APIURL, label, plugin.State, plugin.InstallNum, plugin.Score, plugin.Heat, plugin.CreatedAt, plugin.UpdatedAt)
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

func (m *pluginInfra) CountPlugins(ctx context.Context, fuzzyName ...string) (int, error) {
	var total int
	sql := "SELECT COUNT(*) FROM plugin"
	if len(fuzzyName) != 0 && fuzzyName[0] != "" {
		sql = fmt.Sprintf("SELECT COUNT(*) FROM plugin WHERE name LIKE '%%%s%%'", fuzzyName[0])
	}
	err := m.db.QueryRow(sql).Scan(&total)
	klog.Infof("count plugins succ, total is: %d", total)
	return total, err
}

func (m *pluginInfra) UpdatePluginScoreAndHeat(ctx context.Context, id, heat int32, score float64) error {
	r, err := m.db.ExecContext(ctx, "UPDATE plugin SET score = ?, heat = ? WHERE id = ?", score, heat, id)
	if err != nil {
		klog.Errorf("update plugin score and heat failed, update err: %v", err)
		return err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		klog.Errorf("update plugin score and heat failed, rows_affected err: %v", err)
		return err
	}
	klog.Infof("update plugin score and heat succ, affected: %d", affected)
	return nil
}
