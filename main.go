package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"gin-base/config"
	"gin-base/pkg/global"
	pkg_log "gin-base/pkg/log"
	"gin-base/pkg/rabbitmq"
	"gin-base/pkg/schedule"
	"gin-base/routes"
)

var (
	cfgPath = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	pflag.Parse()

	conf := config.Perform(*cfgPath)
	app := global.Perform(conf)
	defer app.DB.Close()

	gin.SetMode(global.ModeRelease)
	if viper.GetString("project.run_mode") == global.ModeDebug {
		gin.SetMode(global.ModeDebug)
		app.DB.Debug()
	}

	rabbitmq.Perform()

	schedule.Init()

	logConfig := pkg_log.Config{
		Writers:    viper.GetString("log.writers"),
		Level:      viper.GetString("log.level"),
		File:       viper.GetString("log.file"),
		WarnFile:   viper.GetString("log.warn_file"),
		ErrorFile:  viper.GetString("log.error_file"),
		MaxSize:    viper.GetInt("log.max_size"),
		MaxBackups: viper.GetInt("log.max_backups"),
		MaxAge:     viper.GetInt("log.max_age"),
	}
	err := pkg_log.New(&logConfig, pkg_log.InstanceZapLogger)
	if err != nil {
		log.Printf("%s: %v", "Log Init Failed", err)
	}

	router := routes.Perform()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
