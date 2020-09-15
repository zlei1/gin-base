package admin

import (
	"strconv"

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
// @Param page body int true "页"
// @Param per_page body int true "页数"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins [get]
func IndexAdmin(c *gin.Context) {
	var req request.IndexAdminRequest
	_ = c.ShouldBindQuery(&req)

	items, total_count, err := services.AdminSvc.AdminList(&req)
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
// @Router /api/admin/admins/:id [get]
func ShowAdmin(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}

// @Summary 创建管理员
// @Tags admin
// @Produce json
// @Param phone body int true "手机号"
// @Param name body int true "姓名"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins [post]
func CreateAdmin(c *gin.Context) {
	var req request.AdminRequest
	_ = c.Bind(&req)

	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	_, err = services.AdminSvc.AdminCreate(&req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	helpers.SendResponse(c, e.Ok, nil)
}

// @Summary 修改管理员
// @Tags admin
// @Produce json
// @Param phone body int true "手机号"
// @Param name body int true "姓名"
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins/:id [put]
func UpdateAdmin(c *gin.Context) {
	var req request.AdminRequest
	_ = c.Bind(&req)

	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	err = services.AdminSvc.AdminUpdate(uint64(userID), &req)
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	helpers.SendResponse(c, e.Ok, nil)
}

// @Summary 删除管理员
// @Tags admin
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} json "{"code":200,"message":"ok","data":{}}"
// @Router /api/admin/admins/:id [delete]
func DeleteAdmin(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	err := services.AdminSvc.AdminDelete(uint64(userID))
	if err != nil {
		helpers.SendResponse(c, err, nil)
		return
	}

	helpers.SendResponse(c, nil, nil)
}
