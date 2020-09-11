package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"gin-base/app/models"
	"gin-base/config"
	dbRedis "gin-base/pkg/redis"
)

const (
	ModeDebug   string = "debug"
	ModeRelease string = "release"
	ModeTest    string = "test"
)

var (
	App            *Application
	RedisClientNil = redis.Nil
)

type Application struct {
	Conf        *config.Config
	DB          *gorm.DB
	RedisClient *redis.Client
	Debug       bool
}

func Perform(cfg *config.Config) *Application {
	app := new(Application)

	app.Conf = cfg

	app.DB = models.Perform()

	app.RedisClient = dbRedis.Perform()

	if viper.GetString("project.run_mode") == ModeDebug {
		app.Debug = true
	}

	App = app

	return app
}
