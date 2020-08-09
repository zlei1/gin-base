package main

import (
	"github.com/spf13/pflag"

	"gin-base/config"
	"gin-base/routes"
)

var (
	cfgPath = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()

	conf := config.Setup(*cfgPath)
	app := Setup(conf)
	defer app.DB.Close()

	router := routes.Setup()
	router.Run()
}
