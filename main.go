package main

import (
	"github.com/spf13/pflag"

	"gin-base/app/models"
	"gin-base/config"
	"gin-base/pkg/redis"
	"gin-base/routes"
)

var (
	cfgPath = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()

	// init config
	config.Setup(*cfgPath)

	// init log
	config.SetupLog()

	// 连接数据库
	models.Setup()
	defer models.DB.Close()

	// init redis
	redis.Setup()

	// 路由
	router := routes.Setup()
	router.Run()
}
