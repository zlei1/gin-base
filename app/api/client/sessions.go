package client

import (
	"gin-base/app/api/common/helpers"
	"github.com/gin-gonic/gin"
)

// @Summary 管理员登入
// @Produce json
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/client/sessions [post]
func Login(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}
