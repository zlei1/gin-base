package models

import (
	"time"

	"gin-base/pkg/jwt"
)

type Admin struct {
	ID                uint64
	Code              string    `gorm:"column:code" gorm:"comment:'编号'"`
	Name              string    `gorm:"column:name" gorm:"comment:'姓名'"`
	Phone             string    `gorm:"column:phone" gorm:"comment:'手机号'"`
	EncryptedPassword string    `gorm:"column:encrypted_password" gorm:"comment:'密码'"`
	CurrentSignInAt   time.Time `gorm:"column:current_sign_in_at" gorm:"comment:'当前登入时间'"`
	LastSignInAt      time.Time `gorm:"column:last_sign_in_at" gorm:"comment:'上次登入时间'"`
	CurrentSignInIp   string    `gorm:"column:current_sign_in_ip" gorm:"comment:'当前登入Ip'"`
	LastSignInIp      string    `gorm:"column:last_sign_in_ip" gorm:"comment:'上次登入Ip'"`
	Status            int       `gorm:"column:status" gorm:"comment:'状态'"`
	CreatedAt         time.Time `gorm:"column:created_at" gorm:"comment:'创建时间'"`
	UpdatedAt         time.Time `gorm:"column:updated_at" gorm:"comment:'修改时间'"`
}

func (admin *Admin) IssueToken() (token string, err error) {
	j := jwt.JwtContext{
		UserID: admin.ID,
	}

	token, err = jwt.IssueToken(j)

	return
}
