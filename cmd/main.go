package main

import (
	"github.com/GptPluginHub/hub/cmd/app"

	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog"
)

func main() {
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		klog.Error(err)
	}
}
