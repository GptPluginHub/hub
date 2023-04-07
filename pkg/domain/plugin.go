package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/infra/db"
	"github.com/GptPluginHub/hub/pkg/model"

	"hub.io/api/types"
	"k8s.io/klog"
)

const (
	AiPluginURLPath = "/.well-known/ai-plugin.json"
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

func (p *Plugin) ListPluginByFuzzyName(ctx context.Context, fuzzyName, sortFieldName string, orderBy types.OrderBy, page *types.Page) ([]*model.Plugin, error) {
	total, err := p.DBServer.CountPlugins(ctx, fuzzyName)
	if err != nil {
		klog.Errorf("ListPluginByFuzzyName CountPlugins error: %v", err)
		return nil, err
	}

	plugins, err := p.DBServer.ListPlugins(ctx, page.PageSize, page.Page, orderBy.String(), sortFieldName, fuzzyName)
	if err != nil {
		klog.Errorf("ListPluginByFuzzyName ListPlugins error: %v", err)
		return nil, err
	}
	page.Total = int32(total)
	page.Pages = int32(math.Ceil(float64(total) / float64(page.PageSize)))
	return plugins, nil

}

func (p *Plugin) AddPlugin(ctx context.Context, plugin *model.Plugin) error {
	if plugin.CreatedAt == "" {
		plugin.CreatedAt = time.Now().Format(time.RFC3339)
	}
	if plugin.UpdatedAt == "" {
		plugin.UpdatedAt = time.Now().Format(time.RFC3339)
	}
	return p.DBServer.CreatePlugin(ctx, plugin)
}

func (p *Plugin) GeneratePluginURL(ctx context.Context, domain string) string {
	if !strings.HasPrefix(domain, "http") {
		domain = "http://" + domain
	}
	domain = strings.TrimSuffix(domain, "/")
	return domain + AiPluginURLPath
}

func (p *Plugin) FetchAiPluginInfo(aiPluginURL string) (model.AiPlugin, error) {
	resp, err := http.Get(aiPluginURL)
	if err != nil {
		klog.Errorf("FetchAiPluginInfo http.Get error: %v", err)
		return model.AiPlugin{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		klog.Errorf("FetchAiPluginInfo io.ReadAll error: %v", err)
		return model.AiPlugin{}, err
	}
	var aiPlugin model.AiPlugin
	if err = json.Unmarshal(body, &aiPlugin); err != nil {
		klog.Errorf("FetchAiPluginInfo json.Unmarshal error: %v", err)
		return model.AiPlugin{}, err
	}
	return aiPlugin, nil
}

func (p *Plugin) GeneratePluginModel(ctx context.Context, domain, label string, ai model.AiPlugin) model.Plugin {
	plugin := model.Plugin{
		Domain:       domain,
		Name:         ai.NameForHuman,
		Description:  ai.DescriptionForHuman,
		AuthType:     string(ai.Auth.Type),
		LogoURL:      ai.LogoURL,
		ContactEmail: ai.ContactEmail,
		APIType:      ai.API.Type,
		APIURL:       ai.API.URL,
	}
	if label != "" {
		labels := strings.Split(label, ",")
		plugin.Label = labels
	} else {
		plugin.Label = []string{}
	}
	// current use default value.
	plugin.Organization = model.OrganizationTypeTeam.String()
	plugin.State = model.StateTypePublished.String()
	return plugin
}

func (p *Plugin) CheckExist(ctx context.Context, name string) bool {
	plugin, err := p.DBServer.GetPluginByName(ctx, name)
	if err != nil && err != sql.ErrNoRows {
		klog.Errorf("CheckExist GetPluginByName error: %v", err)
		return true
	}
	if err != sql.ErrNoRows {
		return false
	}
	if plugin != nil && plugin.ID != 0 {
		klog.Errorf("CheckExist plugin exist: %v", plugin)
		return true
	}
	return false
}
