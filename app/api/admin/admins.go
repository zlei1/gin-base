package admin

import (
	"github.com/gin-gonic/gin"

	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/api/admin/helpers/response"
	"gin-base/app/api/admin/helpers/services"
	"gin-base/app/api/common/helpers"
	"gin-base/pkg/e"
)

// @Summary 查看管理员
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins [get]
func IndexAdmin(c *gin.Context) {
	var req request.IndexAdminRequest
	_ = c.ShouldBindQuery(&req)

	items, total_count, err := services.GetIndexAdmin(&req)
	if err != nil {
		helpers.SendResponse(c, e.GetIndexAdminError, nil)
	} else {
		helpers.SendResponse(c, e.Ok, response.PaginatorResponse{
			Items:      items,
			TotalCount: total_count,
			Page:       req.Page,
			PerPage:    req.PerPage,
		})
	}
}

// @Summary 查看管理员详情
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins/:code [get]
func ShowAdmin(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}

// @Summary 创建管理员
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins [post]
func CreateAdmin(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}

// @Summary 修改管理员
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins/:code [put]
func UpdateAdmin(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}

// @Summary 删除管理员
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins/:code [delete]
func DeleteAdmin(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}
