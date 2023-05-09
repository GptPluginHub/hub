package domain

import (
	"context"
	"testing"
)

func Test_GeneratePluginMetadata(t *testing.T) {
	pm := PluginMetadata{}
	metadata := pm.GeneratePluginMetadata(context.Background(), 1, "https://www.trip.com/.well-known/ai-plugin.json")
	t.Logf("%+v", metadata)
}
