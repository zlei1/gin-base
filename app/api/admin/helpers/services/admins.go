package services

import (
	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/api/admin/helpers/response"
	"gin-base/app/models"
	"gin-base/pkg/sign"
)

// 管理员登入
func AdminLogin(req *request.AdminLoginRequest) (admin *models.Admin, err error) {
	var a models.Admin

	req.Password = sign.Md5([]byte(req.Password))
	err = models.DB.Where("phone = ? AND encrypted_password = ?", req.Phone, req.Password).First(&a).Error

	return &a, err
}

// 管理员列表
func GetIndexAdmin(req *request.IndexAdminRequest) (list interface{}, total int, err error) {
	if req.PerPage < 1 {
		req.PerPage = 25
	}

	if req.Page < 1 {
		req.Page = 1
	}

	limit := req.PerPage
	offset := req.PerPage * (req.Page - 1)

	db := models.DB.Model(&models.Admin{})
	var items []models.Admin

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&items).Error

	var data []response.IndexAdminResponse
	for _, item := range items {
		data = append(data, response.IndexAdminResponse{
			Code:  item.Code,
			Name:  item.Name,
			Phone: item.Phone,
		})
	}

	return data, total, err
}
