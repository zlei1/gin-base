package admin

import (
	"github.com/gin-gonic/gin"

	"gin-base/api/admin/helpers/request"
	"gin-base/app/api/common/helpers"
)

// @Summary 用户登入
// @Produce json
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/client/sessions [post]
func PhoneLogin(c *gin.Context) {
	var req = request.PhoneLoginRequest
	_ = c.ShouldBindJSON(&req)

	helpers.SendResponse(c, nil, nil)
}
