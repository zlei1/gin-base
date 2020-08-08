package admin

import (
	"github.com/gin-gonic/gin"

	"gin-base/app/api/common/helpers"
)

// @Summary 查看管理员
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins [get]
func GetAdminList(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}
