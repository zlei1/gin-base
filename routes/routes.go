package routes

import (
	"github.com/gin-gonic/gin"

	"gin-base/routes/middleware"
)

func Perform() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Exception())
	r.Use(middleware.RequestId())
	r.Use(middleware.RequestLog())
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	{
		InitAdmin(r)
		InitClient(r)
		InitCommon(r)
		InitSwagger(r)
		InitWebsocket(r)
	}

	return r
}
