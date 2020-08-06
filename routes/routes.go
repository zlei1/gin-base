package routes

import (
	client_api "gin-base/app/api/client"
	common_api "gin-base/app/api/common"
	_ "gin-base/docs"

	"gin-base/routes/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.RequestLog())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// common
	r.GET("/api/common/captcha", common_api.GetCaptcha)
	r.GET("/api/common/phone_verify_code", common_api.GetPhoneVerifyCode)

	// 客户端登入
	// r.GET("/api/client/phone_verify_code", client_api.GetPhoneVerifyCode)
	r.POST("/api/client/sessions", client_api.Login)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
