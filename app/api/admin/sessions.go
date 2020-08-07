package admin

import (
	"github.com/gin-gonic/gin"

	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/api/admin/helpers/response"
	"gin-base/app/api/common/helpers"
	services "gin-base/app/services/admin"
	"gin-base/pkg/e"
)

// @Summary 管理员登入
// @Produce json
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/client/sessions [post]
func Login(c *gin.Context) {
	var req = request.AdminLoginRequest{}
	_ = c.ShouldBindJSON(&req)

	admin, err := services.AdminLogin(&req)
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
