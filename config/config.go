package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
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

	config.initLoging()
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

func (c *Config) initLoging() {
	log_cfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(log_cfg)
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed: ", e.Name)
	})
}
