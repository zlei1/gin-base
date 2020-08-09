package services

import (
	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/models"
)

// 管理员登入
func AdminLogin(req *request.AdminLoginRequest) (admin *models.Admin, err error) {
	var a models.Admin

	err = models.DB.Where("phone = ? AND encrypted_password = ?", req.Phone, req.Password).First(&a).Error

	return &a, err
}

// 管理员列表
func GetIndexAdmin(req *request.IndexAdminRequest) (list interface{}, total int, err error) {
	limit := req.PerPage
	offset := req.PerPage * (req.Page - 1)

	db := models.DB.Model(&models.Admin{})
	var items []models.Admin

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&items).Error
	return items, total, err
}
