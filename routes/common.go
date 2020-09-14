package routes

import (
	"github.com/gin-gonic/gin"

	common_api "gin-base/app/api/common"
)

func InitCommon(router *gin.Engine) {
	r := router.Group("/api/common")
	r.GET("/captcha", common_api.GetCaptcha)
	r.GET("/phone_verify_code", common_api.GetPhoneVerifyCode)
}
