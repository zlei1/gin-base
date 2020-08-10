package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"gin-base/config"
	"gin-base/pkg/global"
	"gin-base/pkg/log"
	"gin-base/routes"
)

var (
	cfgPath = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()

	conf := config.Setup(*cfgPath)
	app := global.Setup(conf)
	defer app.DB.Close()

	logConfig := log.Config{
		Writers:    viper.GetString("log.writers"),
		Level:      viper.GetString("log.level"),
		File:       viper.GetString("log.file"),
		WarnFile:   viper.GetString("log.warn_file"),
		ErrorFile:  viper.GetString("log.error_file"),
		MaxSize:    viper.GetInt("log.max_size"),
		MaxBackups: viper.GetInt("log.max_backups"),
		MaxAge:     viper.GetInt("log.max_age"),
	}
	err := log.New(&logConfig, log.InstanceZapLogger)
	if err != nil {
		fmt.Printf("setup log err: %v", err)
	}

	router := routes.Setup()
	router.Run()
}
