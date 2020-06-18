package routes

import (
	client_api "gin-base/app/api/client"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 客户端登入
	r.GET("/api/client/phone_verify_code", client_api.GetPhoneVerifyCode)
	r.POST("/api/client/sessions", client_api.Login)

	return r
}
