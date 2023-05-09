package application

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/domain"

	"k8s.io/klog"
)

type OpenAPIAppInterface interface {
	// GetOpenAPIFilePath return local cache openapi json file path
	GetOpenAPIFilePath(apiURL string) (string, error)
	// GetOpenAPIData return openapi json data
	GetOpenAPIData(ctx context.Context, apiURL string) (io.ReadCloser, error)
}

var _ OpenAPIAppInterface = new(OpenAPIApp)

type OpenAPIApp struct {
	openAPIDir string
	domain.OpenAPICacheInterface
	domain.Plugin
	domain.PluginMetadata
}

// GetOpenAPIFilePath return local cache openapi json file path.
func (o *OpenAPIApp) GetOpenAPIFilePath(apiURL string) (string, error) {
	key := o.Key(apiURL)
	klog.Info("cache key is: ", key)
	cacheEntry, err := o.generateEntry(apiURL)
	if err != nil {
		klog.Errorf("generate openapi cache entry failed: %v", err)
		return "", err
	}
	if err = o.Set(key, cacheEntry); err != nil {
		return "", err
	}
	return cacheEntry.APIFilePath, nil
}

func (o *OpenAPIApp) GetOpenAPIData(ctx context.Context, apiURL string) (io.ReadCloser, error) {
	plugin, err := o.Plugin.PluginInfra.GetPluginByAPIURL(ctx, apiURL)
	if err != nil {
		klog.Errorf("GetOpenAPIData GetPluginByAPIURL error: %v", err)
		return nil, err
	}
	metaData, err := o.PluginMetadata.PluginMetadataInfra.SelectPluginMetaDataAll(ctx, plugin.ID)
	if err != nil {
		klog.Errorf("GetOpenAPIData SelectPluginMetaDataAll error: %v", err)
		return nil, err
	}
	if metaData.PluginAPI == "" {
		pluginMetadata := o.PluginMetadata.GeneratePluginMetadata(ctx, plugin.ID, apiURL)
		pluginMetadata.ID = metaData.ID
		if err := o.PluginMetadata.UpdatePluginMetadata(ctx, pluginMetadata); err != nil {
			klog.Errorf("GetOpenAPIData UpdatePluginMetadata error: %v", err)
			return nil, errors.New("pluginAPI is empty")
		}
		metaData.PluginAPI = pluginMetadata.PluginAPI
		if metaData.PluginAPI == "" {
			klog.Errorf("GetOpenAPIData pluginAPI is empty")
			return nil, errors.New("pluginAPI is empty")
		}
	}
	return io.NopCloser(strings.NewReader(metaData.PluginAPI)), nil
}

func (o *OpenAPIApp) generateEntry(apiURL string) (domain.OpenAPICacheEntry, error) {
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return domain.OpenAPICacheEntry{}, err
	}
	key := o.Key(apiURL)
	filename := filepath.Base(parsedURL.Path)
	if filename == "" {
		filename = "openapi.json"
	}
	filename = fmt.Sprintf("%s-%s", key, filename)
	resp, err := http.Get(apiURL)
	if err != nil {
		if entry, exist := o.Get(key); exist {
			klog.Warning("get openapi from cache, because get openapi from remote failed")
			return entry, nil
		}
		return domain.OpenAPICacheEntry{}, err
	}

	if o.IsNotExpires(key, resp.Header.Get("Etag"), resp.Header.Get("last-modified")) {
		klog.Info("cache is not expires")
		resp.Body.Close()
		entry, _ := o.Get(key)
		return entry, nil
	}
	defer resp.Body.Close()
	if _, err = os.Stat(o.openAPIDir); os.IsNotExist(err) {
		os.MkdirAll(o.openAPIDir, os.ModePerm)
	}
	filePath := filepath.Join(o.openAPIDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return domain.OpenAPICacheEntry{}, err
	}
	io.Copy(file, resp.Body)
	return domain.OpenAPICacheEntry{
		APIFilePath:  filePath,
		LastModified: resp.Header.Get("last-modified"),
		ETag:         resp.Header.Get("Etag"),
	}, nil
}

func NewOpenAPIApp(dir, fileName string, mysqlConf *config.MysqlOptions) OpenAPIAppInterface {
	if dir == "" {
		dir = domain.DefaultCacheDir
	}
	openAPIDir := filepath.Join(dir, domain.DefaultOpenAPIDirName)
	openAPICache := domain.NewOpenAPICache(dir, fileName)
	plugin, err := domain.NewPlugin(mysqlConf)
	if err != nil {
		panic(err)
	}
	pluginMetadata, err := domain.NewPluginMetadata(mysqlConf)
	if err != nil {
		panic(err)
	}
	return &OpenAPIApp{
		openAPIDir:            openAPIDir,
		OpenAPICacheInterface: openAPICache,
		Plugin:                *plugin,
		PluginMetadata:        *pluginMetadata,
	}
}
