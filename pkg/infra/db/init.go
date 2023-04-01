package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/GptPluginHub/hub/pkg/config"

	"k8s.io/klog"
)

type MysqlServer struct {
	PluginInfraInterface
}

var (
	globalDB *sql.DB
	once     sync.Once
)

func dbInit(mysqlConf *config.MysqlOptions) *sql.DB {
	if globalDB == nil {
		once.Do(func() {
			dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
			klog.Infof("mysql dataSourceName: %s", dataSourceName)
			db, err := sql.Open("mysql", dataSourceName)
			if err != nil {
				panic(err)
			}
			globalDB = db
		})
	}
	return globalDB
}

func NewMysqlServer(mysqlConf *config.MysqlOptions) (*MysqlServer, error) {
	db := dbInit(mysqlConf)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MysqlServer{
		NewPluginInfra(db),
	}, nil
}
