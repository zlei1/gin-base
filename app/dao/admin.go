package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"gin-base/app/models"
)

type AdminBase interface {
	Create(db *gorm.DB, admin models.Admin) (*models.Admin, error)
	Update(db *gorm.DB, id uint64, adminMap map[string]interface{}) error
	Delete(db *gorm.DB, id uint64) error
	GetAdminByID(db *gorm.DB, id uint64) (*models.Admin, error)
	GetAdminByLogin(db *gorm.DB, phone string, password string) (*models.Admin, error)
}

type adminDao struct{}

func NewAdminDao() AdminBase {
	return &adminDao{}
}

// 创建管理员
func (dao *adminDao) Create(db *gorm.DB, admin models.Admin) (*models.Admin, error) {
	err := db.Create(&admin).Error
	if err != nil {
		return nil, errors.Wrap(err, "admin record create err")
	}

	return &admin, nil
}

// 更新管理员
func (dao *adminDao) Update(db *gorm.DB, id uint64, adminMap map[string]interface{}) error {
	admin, err := dao.GetAdminByID(db, id)
	if err != nil {
		return errors.Wrap(err, "admin record update err")
	}

	return db.Model(admin).Update(adminMap).Error
}

// 删除管理员
func (dao *adminDao) Delete(db *gorm.DB, id uint64) error {
	return db.Delete(&models.Admin{}, id).Error
}

// 根据ID获取管理员
func (dao *adminDao) GetAdminByID(db *gorm.DB, id uint64) (*models.Admin, error) {
	admin := models.Admin{}

	err := db.Where(&models.Admin{ID: id}).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "admin record not found")
	}

	return &admin, nil
}

// 根据手机号跟密码获取管理员
func (dao *adminDao) GetAdminByLogin(db *gorm.DB, phone string, password string) (*models.Admin, error) {
	admin := models.Admin{}

	err := db.Where("phone = ? AND encrypted_password = ?", phone, password).First(&admin).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, errors.Wrap(err, "admin login phone or password invalid")
	}

	return &admin, err
}
