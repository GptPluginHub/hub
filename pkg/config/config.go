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
}

type MysqlOptions struct {
	// mysql host
	Host string
	// mysql port
	Port int
	// mysql user
	User string
	// mysql password
	Password string
	// mysql database
	Database string
}

func (m *MysqlOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&m.Host, "mysql-host", m.Host, "mysql host")
	fs.IntVar(&m.Port, "mysql-port", m.Port, "mysql port")
	fs.StringVar(&m.User, "mysql-user", m.User, "mysql user")
	fs.StringVar(&m.Password, "mysql-password", m.Password, "mysql password")
	fs.StringVar(&m.Database, "mysql-database", m.Database, "mysql database")
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

func NewConfig() *Config {
	return &Config{
		Debug:        false,
		MysqlOptions: NewDefaultMysqlOption(),
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
