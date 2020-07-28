package main

import (
	"gin-base/config"
	"gin-base/db"
	"gin-base/pkg/redis"
	"gin-base/routes"

	"github.com/spf13/pflag"
)

var (
	cfgPath = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()

	// init config
	config.Setup(*cfgPath)

	// 连接数据库
	db.OpenDB()
	defer db.GDB.Close()

	// 数据迁移
	// db.MigrateDB()

	// init redis
	redis.Setup()

	// 路由
	router := routes.Setup()
	router.Run()
}
