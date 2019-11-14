package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	ConfigFile string
}

func Init(configFile string) error {
	config := Config{
		ConfigFile: configFile,
	}

	if err := config.init(); err != nil {
		return err
	}

	config.watchConfig()

	return nil
}

func (c *Config) init() error {
	if c.ConfigFile == "" {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(c.ConfigFile)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("LoginServer")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed: ", e.Name)
	})
}
