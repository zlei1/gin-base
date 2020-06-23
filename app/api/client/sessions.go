package client

import (
	"gin-base/app/api/client/helpers"
	"github.com/gin-gonic/gin"

	"gin-base/pkg/e"
)

// 获取手机验证码
// parameter
//   phone
//   code
func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	if phone == "" {
		helpers.SendResponse(c, e.ParamPhoneEmpty, nil)

		return
	}

	code := c.PostForm("code")
	if code == "" {
		helpers.SendResponse(c, e.ParamCodeEmpty, nil)

		return
	}

	helpers.SendResponse(c, nil, nil)
}
