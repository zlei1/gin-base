package client

import (
	"gin-base/app/api/client/helpers"
	"github.com/gin-gonic/gin"

	"gin-base/pkg/e"
)

// @Summary 用户登入
// @Produce json
// @Param phone query string true "Phone"
// @Param code query int false "Code"
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/client/sessions [post]
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
