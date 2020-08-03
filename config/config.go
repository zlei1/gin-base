package config

import (
	"fmt"

	"gin-base/pkg/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Setup(cfgPath string) {
	c := Config{
		Name: cfgPath,
	}

	c.read()
	c.watch()
}

func (cfg *Config) read() {
	if cfg.Name != "" {
		viper.SetConfigFile(cfg.Name)
	} else {
		viper.AddConfigPath("config/")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func (cfg *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func SetupLog() {
	config := log.Config{
		Writers:    viper.GetString("log.writers"),
		Level:      viper.GetString("log.level"),
		File:       viper.GetString("log.file"),
		WarnFile:   viper.GetString("log.warn_file"),
		ErrorFile:  viper.GetString("log.error_file"),
		MaxSize:    viper.GetInt("log.max_size"),
		MaxBackups: viper.GetInt("log.max_backups"),
		MaxAge:     viper.GetInt("log.max_age"),
	}
	err := log.New(&config, log.InstanceZapLogger)
	if err != nil {
		fmt.Printf("setup log err: %v", err)
	}
}
