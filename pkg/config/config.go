package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"k8s.io/klog"
)

const (
	// defaultConfigFile is the config file path.
	defaultConfigFile = "/etc/hub/config.yaml"
)

type Config struct {
	Debug        bool          `json:"debug" yaml:"debug" mapstructure:"debug"`
	MysqlOptions *MysqlOptions `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	APICacheConf *APICache     `json:"api_cache_conf" yaml:"api_cache_conf" mapstructure:"api_cache_conf"`
}

func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&c.Debug, "debug", c.Debug, "debug mode")
	c.MysqlOptions.AddFlags(fs)
	c.APICacheConf.AddFlags(fs)
}

type MysqlOptions struct {
	// mysql host
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	// mysql port
	Port int `json:"port" yaml:"port" mapstructure:"port"`
	// mysql user
	User string `json:"user" yaml:"user" mapstructure:"user"`
	// mysql password
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	// mysql database
	Database string `json:"database" yaml:"database" mapstructure:"database"`
}

func (m *MysqlOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&m.Host, "mysql-host", m.Host, "mysql host")
	fs.IntVar(&m.Port, "mysql-port", m.Port, "mysql port")
	fs.StringVar(&m.User, "mysql-user", m.User, "mysql user")
	fs.StringVar(&m.Password, "mysql-password", m.Password, "mysql password")
	fs.StringVar(&m.Database, "mysql-database", m.Database, "mysql database")
}

type APICache struct {
	Dir      string `json:"dir" yaml:"dir" mapstructure:"dir"`
	FileName string `json:"file_name" yaml:"file_name" mapstructure:"file_name"`
}

func (a *APICache) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&a.Dir, "api-cache-dir", a.Dir, "api cache dir")
	fs.StringVar(&a.FileName, "api-cache-file-name", a.FileName, "api cache file name")
}

func NewDefaultMysqlOption() *MysqlOptions {
	return &MysqlOptions{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "12345678",
		Database: "hub",
	}
}

func NewDefaultAPICacheConf() *APICache {
	return &APICache{
		Dir: "/tmp",
	}
}

func NewConfig() *Config {
	return &Config{
		Debug:        false,
		MysqlOptions: NewDefaultMysqlOption(),
		APICacheConf: NewDefaultAPICacheConf(),
	}
}

func TryLoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		configPath = defaultConfigFile
	}
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		klog.Errorf("try load config: %+v", err)
		return nil, err
	}
	config := NewConfig()
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
