package admin

import (
	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/models"
)

func AdminLogin(req *request.AdminLoginRequest) (admin *models.Admin, err error) {
	var a models.Admin

	err = models.DB.Where("phone = ? AND encrypted_password = ?", req.Phone, req.Password).First(&a).Error

	return &a, err
}
