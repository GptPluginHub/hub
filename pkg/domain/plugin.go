package domain

import (
	"context"
	"encoding/json"
	"io"
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

func (p *Plugin) ListPluginByFuzzyName(ctx context.Context, fuzzyName string, page ...types.Page) ([]*model.Plugin, error) {
	panic("implement me")
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
		labelsJSONBytes, err := json.Marshal(labels)
		if err != nil {
			klog.Errorf("GeneratePluginModel json.Marshal error: %v", err)
		}
		plugin.Label = string(labelsJSONBytes)
	} else {
		plugin.Label = "[]"
	}
	// current use default value.
	plugin.Organization = model.OrganizationTypeTeam.String()
	plugin.State = model.StateTypePublished.String()
	return plugin
}

func (p *Plugin) CheckExist(ctx context.Context, domain string) bool {
	plugin, err := p.DBServer.GetPluginByDomain(ctx, domain)
	if err != nil {
		klog.Errorf("CheckExist GetPluginByDomain error: %v", err)
		return true
	}
	if plugin != nil && plugin.ID != 0 {
		klog.Errorf("CheckExist plugin exist: %v", plugin)
		return true
	}
	return false
}
