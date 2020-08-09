package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"gin-base/app/models"
	"gin-base/config"
	"gin-base/pkg/log"
	pkg_redis "gin-base/pkg/redis"
)

const (
	ModeDebug   string = "debug"
	ModeRelease string = "release"
	ModeTest    string = "test"
)

var App *Application

type Application struct {
	Conf        *config.Config
	DB          *gorm.DB
	RedisClient *redis.Client
	Debug       bool
}

func Setup(cfg *config.Config) *Application {
	app := new(Application)

	app.DB = models.Setup()

	app.RedisClient = pkg_redis.Setup()

	SetupLog()

	if viper.GetString("run_mode") == ModeDebug {
		app.Debug = true
	}

	return app
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
