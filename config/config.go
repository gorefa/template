package config

import (
	"strings"

	"gogin/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name   string
	Logger *logger.Logger
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// init config file
	configFile, err := c.initConfig()
	if err != nil {
		return err
	}

	// init log
	if err := c.initLog(); err != nil {
		return err
	}
	c.Logger = logger.GetLogger("config")
	c.Logger.Debug("config file used", logger.String("file", configFile))

	c.watchConfig()

	return nil
}

func (c *Config) initLog() error {
	viper.SetDefault("log.level", "debug")
	logConf := logger.Config{Level: viper.GetString("log.level")}
	return logger.InitLogger(logConf)
}

func (c *Config) initConfig() (string, error) {
	if c.Name != "" {
		viper.SetConfigName(c.Name)
	} else {
		viper.AddConfigPath("config")
		viper.SetConfigName("cfg")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GOGIN")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}
	return viper.GetViper().ConfigFileUsed(), nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		c.Logger.Info("config file changed", logger.String("file", e.Name))
	})
}
