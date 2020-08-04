package routes

import (
	client_api "gin-base/app/api/client"
	common_api "gin-base/app/api/common"
	_ "gin-base/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// common
	r.GET("/api/common/captcha", common_api.GetCaptcha)

	// 客户端登入
	r.GET("/api/client/phone_verify_code", client_api.GetPhoneVerifyCode)
	r.POST("/api/client/sessions", client_api.Login)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
