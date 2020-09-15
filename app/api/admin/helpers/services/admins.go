package services

import (
	"github.com/jinzhu/gorm"

	"gin-base/app/api/admin/helpers/request"
	"gin-base/app/api/admin/helpers/response"
	"gin-base/app/dao"
	"gin-base/app/models"
	"gin-base/pkg/e"
	"gin-base/pkg/global"
	"gin-base/pkg/sign"
)

type AdminServer interface {
	LoginAdminFind(req *request.AdminRequest) (admin *models.Admin, err error)
	AdminList(req *request.IndexAdminRequest) (list interface{}, total int, err error)
	AdminCreate(req *request.AdminRequest) (admin *models.Admin, err error)
	AdminUpdate(id uint64, req *request.AdminRequest) (err error)
	AdminDelete(id uint64) (err error)
}

var AdminSvc = NewAdminService()

type adminServer struct {
	adminDao dao.AdminBase
}

func NewAdminService() *adminServer {
	return &adminServer{
		adminDao: dao.NewAdminDao(),
	}
}

// 管理员登入
func (server *adminServer) LoginAdminFind(req *request.AdminLoginRequest) (admin *models.Admin, err error) {
	admin, err = server.adminDao.GetAdminByLogin(
		global.App.DB,
		req.Phone,
		sign.Md5([]byte(req.Password)),
	)

	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, e.AdminLoginError
	}

	return
}

// 管理员列表
func (server *adminServer) AdminList(req *request.IndexAdminRequest) (list interface{}, total int, err error) {
	if req.PerPage < 1 {
		req.PerPage = 25
	}

	if req.Page < 1 {
		req.Page = 1
	}

	limit := req.PerPage
	offset := req.PerPage * (req.Page - 1)

	db := global.App.DB.Model(&models.Admin{})
	var items []models.Admin

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&items).Error

	var data []response.IndexAdminResponse
	for _, item := range items {
		data = append(data, response.IndexAdminResponse{
			Name:  item.Name,
			Phone: item.Phone,
		})
	}

	return data, total, err
}

// 创建管理员
func (server *adminServer) AdminCreate(req *request.AdminRequest) (admin *models.Admin, err error) {
	password := sign.Md5([]byte(req.Phone))

	admin, err = server.adminDao.Create(
		global.App.DB,
		models.Admin{
			Name:              req.Name,
			Phone:             req.Phone,
			EncryptedPassword: password,
		},
	)

	return
}

// 修改管理员
func (server *adminServer) AdminUpdate(id uint64, req *request.AdminRequest) (err error) {
	adminMap := make(map[string]interface{})
	adminMap["phone"] = req.Phone
	adminMap["name"] = req.Name

	err = server.adminDao.Update(global.App.DB, id, adminMap)

	return
}

// 删除管理员
func (server *adminServer) AdminDelete(id uint64) (err error) {
	err = server.adminDao.Delete(global.App.DB, id)

	return
}
