package model

// PluginScore is the model for plugin_score table.
type PluginScore struct {
	ID        int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	PluginID  int    `gorm:"column:plugin_id" db:"plugin_id" json:"plugin_id" form:"plugin_id"`
	Score     int    `gorm:"column:score" db:"score" json:"score" form:"score"`
	Comments  string `gorm:"column:comments" db:"comments" json:"comments" form:"comments"`
	UserID    int    `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	CreatedAt string `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

// Rating is computed rating for a plugin.
type Rating struct {
	Value float64
	Count int
}
