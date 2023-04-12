package domain

import (
	"context"
	"time"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/infra/db"
	"github.com/GptPluginHub/hub/pkg/model"
)

type PluginScore struct {
	PluginScoreInfra db.PluginScoreInfraInterface
}

func (p *PluginScore) AddPluginScore(ctx context.Context, score model.PluginScore) error {
	if score.CreatedAt == "" {
		score.CreatedAt = time.Now().Format(time.RFC3339)
	}
	if score.UpdatedAt == "" {
		score.UpdatedAt = time.Now().Format(time.RFC3339)
	}
	return p.PluginScoreInfra.InsertPluginScore(ctx, score)
}

func (p *PluginScore) ListPluginScoreByID(ctx context.Context, pluginID int) ([]*model.PluginScore, error) {
	if pluginID == 0 {
		return nil, nil
	}
	return p.PluginScoreInfra.SelectPluginScore(ctx, pluginID)
}

func (p *PluginScore) CalculateRating(ctx context.Context, scores []float64) model.Rating {
	rating := model.Rating{Value: 0, Count: 0}
	for _, score := range scores {
		rating.Value = (rating.Value*float64(rating.Count) + score) / float64(rating.Count+1)
		rating.Count++
	}
	return rating
}

func NewPluginScore(mysqlConf *config.MysqlOptions) (*PluginScore, error) {
	pluginScoreInfra, err := db.NewPluginScoreInfra(mysqlConf)
	if err != nil {
		return nil, err
	}
	return &PluginScore{PluginScoreInfra: pluginScoreInfra}, nil
}
