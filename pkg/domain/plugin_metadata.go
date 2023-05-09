package domain

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/infra/db"
	"github.com/GptPluginHub/hub/pkg/model"

	"k8s.io/klog"
)

type PluginMetadata struct {
	PluginMetadataInfra db.PluginMetadataInfraInterface
}

func (p *PluginMetadata) AddPluginMetadata(ctx context.Context, metadata model.PluginMetadata) error {
	return p.PluginMetadataInfra.InsertPluginMetadata(ctx, metadata)
}

func (p *PluginMetadata) UpdatePluginMetadata(ctx context.Context, metadata model.PluginMetadata) error {
	return p.PluginMetadataInfra.UpdatePluginMetadata(ctx, metadata)
}

func (p *PluginMetadata) GeneratePluginMetadata(ctx context.Context, pluginID int, pluginURL string) model.PluginMetadata {
	schemaData, openAPIURL, header, err := p.getPluginSchemaData(ctx, pluginURL)
	if err != nil {
		klog.Errorf("GeneratePluginMetadata getPluginSchemaData error: %v", err)
		return model.PluginMetadata{}
	}
	openAPIData, err := p.getOpenAPIData(ctx, openAPIURL)
	if err != nil {
		klog.Errorf("GeneratePluginMetadata getOpenAPIData error: %v", err)
		return model.PluginMetadata{}
	}
	return model.PluginMetadata{
		PluginID:              pluginID,
		PluginAPIEtag:         header.Get("Etag"),
		PluginAPILastModified: header.Get("last-modified"),
		PluginSchema:          schemaData,
		PluginAPI:             openAPIData,
		CreateAt:              time.Now().Format(time.RFC3339),
		UpdateAt:              time.Now().Format(time.RFC3339),
	}
}

func (p *PluginMetadata) getPluginSchemaData(ctx context.Context, pluginURL string) (string, string, http.Header, error) {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, pluginURL, nil)
	schemaResp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", "", nil, err
	}
	defer schemaResp.Body.Close()
	if schemaResp.StatusCode != http.StatusOK {
		return "", "", nil, err
	}
	readAll, err := io.ReadAll(schemaResp.Body)
	if err != nil {
		return "", "", nil, err
	}
	var aiPlugin model.AiPlugin
	if err = json.Unmarshal(readAll, &aiPlugin); err != nil {
		return "", "", nil, err
	}
	return string(readAll), aiPlugin.API.URL, schemaResp.Header, nil
}

func (p *PluginMetadata) getOpenAPIData(ctx context.Context, openAPIURL string) (string, error) {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, openAPIURL, nil)
	openAPIResp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer openAPIResp.Body.Close()
	if openAPIResp.StatusCode != http.StatusOK {
		return "", err
	}
	readAll, err := io.ReadAll(openAPIResp.Body)
	if err != nil {
		return "", err
	}
	return string(readAll), nil
}

func NewPluginMetadata(mysqlConf *config.MysqlOptions) (*PluginMetadata, error) {
	pluginMetadataInfra, err := db.NewPluginMetadataInfra(mysqlConf)
	if err != nil {
		return nil, err
	}
	return &PluginMetadata{PluginMetadataInfra: pluginMetadataInfra}, nil
}
