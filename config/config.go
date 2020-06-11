package config

import (
	"fmt"
	"log"

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
		viper.SetConfigName("config.local")
	}
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
