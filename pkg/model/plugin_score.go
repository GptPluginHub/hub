package model

type PluginScore struct {
	ID        int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	PluginID  int    `gorm:"column:plugin_id" db:"plugin_id" json:"plugin_id" form:"plugin_id"`
	Score     int    `gorm:"column:score" db:"score" json:"score" form:"score"`
	Comments  string `gorm:"column:comments" db:"comments" json:"comments" form:"comments"`
	UserID    int    `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	CreatedAt int64  `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}
