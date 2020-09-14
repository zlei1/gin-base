package routes

import (
	"github.com/gin-gonic/gin"

	client_api "gin-base/app/api/client"
	"gin-base/routes/middleware"
)

func InitClient(router *gin.Engine) {
	r := router.Group("/api/client")
	r.POST("/sessions", client_api.Login)

	auth := r.Group("")
	auth.Use(middleware.Auth())
}
