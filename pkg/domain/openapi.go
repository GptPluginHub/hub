package domain

import (
	"encoding/json"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"k8s.io/klog"
)

const (
	// DefaultCacheDir is the default directory for the cache dir.
	DefaultCacheDir = "."
	// DefaultCacheFileName is the default name of the cache file.
	DefaultCacheFileName = "openapi_cache.json"
	// DefaultOpenAPIDirName is the default name of the openapi file dir.
	DefaultOpenAPIDirName = "openapi"
)

type OpenAPICacheInterface interface {
	// Get returns the OpenAPICacheEntry for the given url if it exists in the cache.
	Get(key string) (OpenAPICacheEntry, bool)
	// Set adds or updates the given entry for the given url in the cache.
	Set(key string, entry OpenAPICacheEntry) error
	// IsNotExpires check cache is expires
	IsNotExpires(key, etag, lastModified string) bool
	// Key returns the cache key for the given apiUrl.
	Key(apiURL string) string
}

var _ OpenAPICacheInterface = new(openAPICache)

type openAPICache struct {
	dir      string
	fileName string
	data     map[string]OpenAPICacheEntry
	sync.RWMutex
}

type OpenAPICacheEntry struct {
	ETag         string `json:"eTag,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	APIFilePath  string `json:"apiFileName,omitempty"`
}

func NewOpenAPICache(dir, fileName string) OpenAPICacheInterface {
	if dir == "" {
		dir = DefaultCacheDir
	}
	if fileName == "" {
		fileName = DefaultCacheFileName
	}
	filePath := filepath.Join(dir, fileName)
	file, err := os.Open(filePath)
	if err != nil && !os.IsNotExist(err) {
		klog.Errorf("Failed to open cache file %s: %v", filepath.Join(dir, fileName), err)
		panic(err)
	}
	if os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			klog.Errorf("Failed to create cache file %s: %v", filepath.Join(dir, fileName), err)
			panic(err)
		}
	}
	defer file.Close()
	var data map[string]OpenAPICacheEntry
	bytes, err := io.ReadAll(file)
	if err != nil {
		if err = os.WriteFile(filePath, []byte("{}"), 0o644); err != nil {
			panic(err)
		}
		klog.Warningf("Failed to read cache file %s: %v, write empty object to file", filePath, err)
	} else {
		klog.Infof("OpenAPI cache file %s: %s", filePath, string(bytes))
		if err = json.Unmarshal(bytes, &data); err != nil {
			klog.Errorf("Failed to unmarshal cache file %s: %v", filePath, err)
			panic(err)
		}
	}
	if data == nil {
		data = make(map[string]OpenAPICacheEntry)
	}
	return &openAPICache{
		dir:      dir,
		fileName: fileName,
		data:     data,
	}
}

func (o *openAPICache) Get(key string) (OpenAPICacheEntry, bool) {
	o.RLock()
	defer o.RUnlock()
	entry, ok := o.data[key]
	return entry, ok
}

func (o *openAPICache) Set(key string, entry OpenAPICacheEntry) error {
	if o.IsNotExpires(key, entry.ETag, entry.LastModified) {
		klog.Infof("OpenAPI cache is not expired, key: %s, entry: %v", key, entry)
		return nil
	}
	o.Lock()
	defer o.Unlock()
	o.data[key] = entry
	filePath := filepath.Join(o.dir, o.fileName)
	bytes, err := json.Marshal(o.data)
	if err != nil {
		return err
	}
	if err = os.WriteFile(filePath, bytes, 0o644); err != nil {
		return err
	}
	return nil
}

func (o *openAPICache) Key(apiURL string) string {
	parse, err := url.Parse(apiURL)
	if err != nil {
		klog.Errorf("Failed to parse url %s: %v", apiURL, err)
		return apiURL
	}
	return strings.ReplaceAll(parse.Hostname(), ".", "-")
}

func (o *openAPICache) IsNotExpires(key, etag, lastModified string) bool {
	o.RLock()
	defer o.RUnlock()
	if e, ok := o.data[key]; ok {
		return e.ETag == etag && e.LastModified == lastModified
	}
	return false
}
