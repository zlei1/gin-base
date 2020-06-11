package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"

	"gin-base/config"
	"gin-base/db"
	"gin-base/pkg/redis"
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
	db.MigrateDB()

	// init redis
	redis.Setup()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
