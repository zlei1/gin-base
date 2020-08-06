package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	admin_api "gin-base/app/api/admin"
	client_api "gin-base/app/api/client"
	common_api "gin-base/app/api/common"
	_ "gin-base/docs"
	"gin-base/routes/middleware"
	"github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.RequestLog())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	auth := r.Group("")
	auth.Use(middleware.Auth())

	// common 公共
	r.GET("/api/common/captcha", common_api.GetCaptcha)
	r.GET("/api/common/phone_verify_code", common_api.GetPhoneVerifyCode)

	// admin 管理端
	auth.GET("/api/admin/admins", admin_api.GetAdminList)

	// client 客户端
	r.POST("/api/client/sessions", client_api.Login)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
