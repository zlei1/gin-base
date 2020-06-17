package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取手机验证码
func GetPhoneVerifyCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
