package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

	"gin-base/app/api/common/helpers"
	"gin-base/app/api/common/helpers/response"
	"gin-base/pkg/e"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		80,
		240,
		6,
		0.7,
		80,
	)
	cp := base64Captcha.NewCaptcha(driver, store)
	token, captcha, err := cp.Generate()
	if err != nil {
		helpers.SendResponse(c, e.CaptchaGenError, nil)

		return
	} else {
		helpers.SendResponse(c, e.Ok, response.CaptchaResponse{
			CaptchaToken: token,
			Captcha:      captcha,
		})
	}
}
