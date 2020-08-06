package admin

import (
	"github.com/gin-gonic/gin"

	"gin-base/app/api/common/helpers"
)

// @Summary 查看管理员
// @Produce json
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/client/sessions [post]
func GetAdminList(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}