package application

import (
	"context"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/domain"
	"github.com/GptPluginHub/hub/pkg/model"

	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
	"k8s.io/klog"
)

type PluginScoreAppInterface interface {
	CreatePluginScore(ctx context.Context, req *pluginv1alpha1.CreatePluginScoreRequest) error
}

var _ PluginScoreAppInterface = new(PluginScoreApp)

type PluginScoreApp struct {
	PluginScore *domain.PluginScore
	Plugin      *domain.Plugin
}

func (p *PluginScoreApp) CreatePluginScore(ctx context.Context, req *pluginv1alpha1.CreatePluginScoreRequest) error {
	// add plugin score for this plugin_id
	err := p.PluginScore.AddPluginScore(ctx, model.PluginScore{
		PluginID: int(req.PluginId),
		Score:    float64(req.Score),
		Comments: req.Comments,
	})
	if err != nil {
		klog.Error("CreatePluginScore error: ", err)
		return err
	}
	// get all plugin score for this plugin_id
	pluginScores, err := p.PluginScore.ListPluginScoreByID(ctx, int(req.PluginId))
	if err != nil {
		klog.Error(err)
		return err
	}

	// count all plugin score for this plugin_id
	scores := make([]float64, len(pluginScores))
	for index, score := range pluginScores {
		scores[index] = score.Score
	}

	// computer score this plugin
	pluginScore := p.PluginScore.CalculateRating(ctx, scores)

	// update plugin score and heat
	return p.Plugin.UpdatePluginScoreAndHeat(ctx, int32(req.PluginId), int32(pluginScore.Count), pluginScore.Value)
}

func NewPluginScoreAppInterface(cfg config.Config) PluginScoreAppInterface {
	pluginScore, err := domain.NewPluginScore(cfg.MysqlOptions)
	if err != nil {
		panic(err)
	}
	plugin, err := domain.NewPlugin(cfg.MysqlOptions)
	if err != nil {
		panic(err)
	}
	return &PluginScoreApp{
		PluginScore: pluginScore,
		Plugin:      plugin,
	}
}
