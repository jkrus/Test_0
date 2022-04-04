package config

import (
	"log"
	"path/filepath"
	"time"

	"github.com/jkrus/kit/config"
	"github.com/jkrus/kit/files"
	"github.com/pkg/errors"
)

type (
	// Config represents the main app's configuration.
	Config struct {
		Host                 string               `yaml:"host"`
		HTTP                 HTTP                 `yaml:"http"`
		DB                   DB                   `yaml:"db"`
		NatsSubscribeOptions NatsSubscribeOptions `yaml:"natsSubscribeOptions"`
		Cache                Cache                `yaml:"cache"`
	}
)

// NewConfig constructs an empty configuration.
func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() error {
	filePath := filepath.Join(files.OsAppRootPath(AppRootPathName, AppName, AppUsage, AppVersion), yamlFileName)
	if files.IsFileExist(filePath) {
		log.Println("Read data from config file in path:", filePath)
		if err := files.ReadFromYamlFile(filePath, c); err != nil {
			return errors.Wrap(err, "Init: read config file filed")
		}
	} else {
		log.Println("Create default config file in path:", filePath)
	}

	c.setDefault()
	if err := files.MakeDirs(filePath); err != nil {
		return errors.Wrap(err, "Init: can not create dirs")
	}
	if err := files.WriteToYamlFile(filePath, c); err != nil {
		return errors.Wrap(err, "Init: create config file filed")
	}

	return nil
}

func (c *Config) Load() error {
	err := config.Load(AppRootPathName, AppName, AppUsage, AppVersion, yamlFileName, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) setDefault() {
	c.Host = "127.0.0.1"

	c.HTTP.Port = 8080

	c.DB.User = "root"
	c.DB.Pass = "secret"
	c.DB.Host = "localhost"
	c.DB.Port = 5432
	c.DB.Name = "postgres"
	c.DB.SSLMode = "disable"

	c.NatsSubscribeOptions.AllowReconnect = true
	c.NatsSubscribeOptions.MaxReconnect = 10
	c.NatsSubscribeOptions.ReconnectWait = 5 * time.Second
	c.NatsSubscribeOptions.Timeout = 1 * time.Second

	c.Cache.Size = 10485760
	c.Cache.RecoverySize = 1000
}
