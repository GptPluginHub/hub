package domain

import (
	"context"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/infra/db"
	"github.com/GptPluginHub/hub/pkg/model"

	"hub.io/api/types"
)

type Plugin struct {
	DBServer *db.MysqlServer
}

func NewPlugin(mysqlConf *config.MysqlOptions) (*Plugin, error) {
	mysqlServer, err := db.NewMysqlServer(mysqlConf)
	if err != nil {
		return nil, err
	}
	return &Plugin{DBServer: mysqlServer}, nil
}

func (p *Plugin) ListPluginByFuzzyName(ctx context.Context, fuzzyName string, page ...types.Page) ([]*model.Plugin, error) {
	panic("implement me")
}
