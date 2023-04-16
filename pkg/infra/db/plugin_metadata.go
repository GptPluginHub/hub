package db

import (
	"context"
	"database/sql"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/model"
)

type PluginMetadataInfraInterface interface {
	InsertPluginMetadata(ctx context.Context, metadata model.PluginMetadata) error
	SelectPluginMetaDataBase(ctx context.Context, pluginID int) (*model.PluginMetadata, error)
	SelectPluginMetaDataAll(ctx context.Context, pluginID int) (*model.PluginMetadata, error)
	UpdatePluginMetadata(ctx context.Context, metadata model.PluginMetadata) error
}

var _ PluginMetadataInfraInterface = new(pluginMetadataInfra)

type pluginMetadataInfra struct {
	db *sql.DB
}

func NewPluginMetadataInfra(mysqlConf *config.MysqlOptions) (PluginMetadataInfraInterface, error) {
	db := dbInit(mysqlConf)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &pluginMetadataInfra{db: db}, nil
}

func (pm *pluginMetadataInfra) InsertPluginMetadata(ctx context.Context, metadata model.PluginMetadata) error {
	insert := `INSERT INTO plugin_metadata (plugin_id, plugin_api_etag, plugin_api_last_modified, plugin_schema, plugin_api, create_at, update_at) VALUES (?,?,?,?,?,?,?)`
	_, err := pm.db.ExecContext(ctx, insert, metadata.PluginID, metadata.PluginAPIEtag, metadata.PluginAPILastModified, metadata.PluginSchema, metadata.PluginAPI, metadata.CreateAt, metadata.UpdateAt)
	return err
}

func (pm *pluginMetadataInfra) SelectPluginMetaDataBase(ctx context.Context, pluginID int) (*model.PluginMetadata, error) {
	selectSQL := `SELECT id, plugin_id, plugin_api_etag, plugin_api_last_modified, create_at, update_at FROM plugin_metadata WHERE plugin_id = ?`
	row := pm.db.QueryRowContext(ctx, selectSQL, pluginID)
	var metadata model.PluginMetadata
	if err := row.Scan(&metadata.ID, &metadata.PluginID, &metadata.PluginAPIEtag, &metadata.PluginAPILastModified, &metadata.PluginSchema, &metadata.PluginAPI, &metadata.CreateAt, &metadata.UpdateAt); err != nil {
		return nil, err
	}
	return &metadata, nil
}

func (pm *pluginMetadataInfra) SelectPluginMetaDataAll(ctx context.Context, pluginID int) (*model.PluginMetadata, error) {
	selectSQL := `SELECT id, plugin_id, plugin_api_etag, plugin_api_last_modified, plugin_schema, plugin_api, create_at, update_at FROM plugin_metadata WHERE plugin_id = ?`
	row := pm.db.QueryRowContext(ctx, selectSQL, pluginID)
	var metadata model.PluginMetadata
	if err := row.Scan(&metadata.ID, &metadata.PluginID, &metadata.PluginAPIEtag, &metadata.PluginAPILastModified, &metadata.PluginSchema, &metadata.PluginAPI, &metadata.CreateAt, &metadata.UpdateAt); err != nil {
		return nil, err
	}
	return &metadata, nil
}

func (pm *pluginMetadataInfra) UpdatePluginMetadata(ctx context.Context, metadata model.PluginMetadata) error {
	update := `UPDATE plugin_metadata SET plugin_api_etag=?, plugin_api_last_modified=?, plugin_schema=?, plugin_api=?, update_at=? WHERE id=?`
	_, err := pm.db.ExecContext(ctx, update, metadata.PluginAPIEtag, metadata.PluginAPILastModified, metadata.PluginSchema, metadata.PluginAPI, metadata.UpdateAt, metadata.ID)
	return err
}
