package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"k8s.io/klog"
)

const (
	host            = "http://127.0.0.1:8000"
	createPluginURL = host + "/apis/hub.io/v1alpha1/plugin"
)

func main() {
	dir, _ := os.Getwd()
	filePath := dir + "/example/data.json"
	klog.Infof(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileDataBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var data []map[string]interface{}
	if err = json.Unmarshal(fileDataBytes, &data); err != nil {
		panic(err)
	}
	for _, item := range data {
		itemBytes, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		reader := bytes.NewReader(itemBytes)
		resp, err := http.Post(createPluginURL, "application/json", reader)
		if err != nil {
			panic(err)
		}
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		klog.Infof("Create plugin response: %s", string(respBytes))
	}
}
