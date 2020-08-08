package common

import (
	"github.com/gin-gonic/gin"

	"gin-base/app/api/common/helpers"
	"gin-base/app/api/common/helpers/request"
	"gin-base/app/services"
	"gin-base/pkg/e"
)

// @Summary 获取手机验证码
// @Produce application/json
// @Param phone query string true "手机号"
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/common/phone_verify_code [get]
func GetPhoneVerifyCode(c *gin.Context) {
	var param request.PhoneVerifyCodeRequest
	_ = c.ShouldBindJSON(&param)

	_, err := services.GetVcode(param.Phone)
	if err != nil {
		helpers.SendResponse(c, e.VcodeGetfrequent, nil)
		return
	}

	_, err = services.GenVcode(param.Phone)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	// 短信发送

	helpers.SendResponse(c, e.Ok, nil)
}
