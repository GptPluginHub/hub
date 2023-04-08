package db

import (
	"context"
	"database/sql"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/model"

	"k8s.io/klog"
)

type PluginScoreInfraInterface interface {
	InsertPluginScore(ctx context.Context, score model.PluginScore) error
	SelectPluginScore(ctx context.Context, pluginID int) ([]*model.PluginScore, error)
}

var _ PluginScoreInfraInterface = new(pluginScoreInfra)

type pluginScoreInfra struct {
	db *sql.DB
}

func (ps *pluginScoreInfra) InsertPluginScore(ctx context.Context, score model.PluginScore) error {
	insert := `INSERT INTO plugin_score (plugin_id, score,comments,user_id, created_at, updated_at) VALUES (?,?,?, ?, ?, ?)`
	_, err := ps.db.ExecContext(ctx, insert, score.PluginID, score.Score, score.Comments, score.UserID, score.CreatedAt, score.UpdatedAt)
	return err
}

func (ps *pluginScoreInfra) SelectPluginScore(ctx context.Context, pluginID int) ([]*model.PluginScore, error) {
	selectSQL := `SELECT id, plugin_id, score, comments, user_id, created_at, updated_at FROM plugin_score WHERE plugin_id = ?`
	rows, err := ps.db.QueryContext(ctx, selectSQL, pluginID)
	if err != nil {
		klog.Error("SelectPluginScore error: ", err)
		return nil, err
	}
	defer rows.Close()
	var scores []*model.PluginScore
	for rows.Next() {
		var score model.PluginScore
		if err := rows.Scan(&score.ID, &score.PluginID, &score.Score, &score.Comments, &score.UserID, &score.CreatedAt, &score.UpdatedAt); err != nil {
			return nil, err
		}
		scores = append(scores, &score)
	}
	return scores, nil
}

func NewPluginScoreInfra(mysqlConf *config.MysqlOptions) (PluginScoreInfraInterface, error) {
	db := dbInit(mysqlConf)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &pluginScoreInfra{db: db}, nil
}
