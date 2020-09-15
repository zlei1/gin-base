package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gin-base/app/api/common/helpers"
	"gin-base/app/api/common/helpers/request"
	"gin-base/app/api/common/helpers/services"
	"gin-base/app/workers"
	"gin-base/pkg/e"
	"gin-base/pkg/rabbitmq"
)

var validate = validator.New()

// @Summary 获取手机验证码
// @Tags common
// @Produce application/json
// @Param phone query string true "手机号"
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/common/phone_verify_code [get]
func GetPhoneVerifyCode(c *gin.Context) {
	var req request.PhoneVerifyCodeRequest
	_ = c.Bind(&req)

	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	vcode, err := services.PhoneVcodeSvc.GetVcode(req.Phone)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	if vcode != "" {
		helpers.SendResponse(c, e.VcodeGetfrequent, nil)
		return
	}

	vcode, err = services.PhoneVcodeSvc.GenVcode(req.Phone)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	mqPublishBody := &rabbitmq.Message{
		Payload: workers.SendSmsParams{
			Worker: "SendSms",
			Phone:  req.Phone,
			Vcode:  vcode,
		},
	}
	rabbitmq.MqPublish(mqPublishBody)

	helpers.SendResponse(c, e.Ok, nil)
}
