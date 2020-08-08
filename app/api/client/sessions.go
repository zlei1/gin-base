package client

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gin-base/app/api/client/helpers/request"
	"gin-base/app/api/client/helpers/response"
	"gin-base/app/api/common/helpers"
	user_service "gin-base/app/services/user"
	vocde_service "gin-base/app/services/vcode"
	"gin-base/pkg/e"
)

var validate = validator.New()

// @Summary 客户端用户登入
// @Tags client
// @Produce json
// @Param phone body string true "手机号"
// @Param vcode body string true "手机验证码"
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/client/sessions [post]
func Login(c *gin.Context) {
	var req = request.UserLoginRequest{}
	_ = c.ShouldBindJSON(&req)

	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	if checkVcode := vocde_service.CheckVcode(req.Phone, req.Vcode); checkVcode {
		helpers.SendResponse(c, e.CaptchaInvalid, nil)
		return
	}

	user, err := user_service.UserLogin(&req)
	if err != nil {
		helpers.SendResponse(c, e.UserLoginError, nil)
		return
	}

	token, err := user.IssueToken()
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	helpers.SendResponse(c, e.Ok, response.Token{
		Token: token,
	})
}
