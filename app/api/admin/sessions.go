package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/api/admin/helpers/response"
	"gin-base/app/api/common/helpers"
	admin_service "gin-base/app/services/admin"
	"gin-base/pkg/e"
)

var validate = validator.New()

// @Summary 管理员登入
// @Tags admin
// @Produce json
// @Param phone body string true "手机号"
// @Param password body string true "密码"
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/sessions [post]
func Login(c *gin.Context) {
	var req = request.AdminLoginRequest{}
	_ = c.ShouldBindJSON(&req)

	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	admin, err := admin_service.AdminLogin(&req)
	if err != nil {
		helpers.SendResponse(c, e.AdminLoginError, nil)
		return
	}

	token, err := admin.IssueToken()
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	helpers.SendResponse(c, e.Ok, response.Token{
		Token: token,
	})
}
