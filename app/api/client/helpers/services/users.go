package services

import (
	"gin-base/app/api/client/helpers/request"
	"gin-base/app/models"
	"gin-base/pkg/global"
)

func UserLogin(req *request.UserLoginRequest) (user *models.User, err error) {
	var u models.User

	err = global.App.DB.Where("phone = ?", req.Phone).First(&u).Error

	return &u, err
}
