package model

type PluginMetadata struct {
	ID                    int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	PluginID              int    `gorm:"column:plugin_id" db:"plugin_id" json:"plugin_id" form:"plugin_id"`
	PluginAPIEtag         string `gorm:"column:plugin_api_etag" db:"plugin_api_etag" json:"plugin_api_etag" form:"plugin_api_etag"`
	PluginAPILastModified string `gorm:"column:plugin_api_last_modified" db:"plugin_api_last_modified" json:"plugin_api_last_modified" form:"plugin_api_last_modified"`
	PluginSchema          string `gorm:"column:plugin_schema" db:"plugin_schema" json:"plugin_schema" form:"plugin_schema"`
	PluginAPI             string `gorm:"column:plugin_api" db:"plugin_api" json:"plugin_api" form:"plugin_api"`
	CreateAt              string `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt              string `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}
