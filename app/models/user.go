package models

import (
	"time"
)

type Admin struct {
	ID                uint      `gorm:"primary_key"`
	Code              string    `gorm:"code" gorm:"comment:'编号'"`
	Name              string    `gorm:"name" gorm:"comment:'姓名'"`
	Phone             string    `gorm:"phone" gorm:"comment:'手机号'"`
	EncryptedPassword string    `gorm:"encrypted_password" gorm:"comment:'密码'"`
	CurrentSignInAt   time.Time `gorm:"current_sign_in_at" gorm:"comment:'当前登入时间'"`
	LastSignInAt      time.Time `gorm:"last_sign_in_at" gorm:"comment:'上次登入时间'"`
	CurrentSignInIp   string    `gorm:"current_sign_in_ip" gorm:"comment:'当前登入Ip'"`
	LastSignInIp      string    `gorm:"last_sign_in_ip" gorm:"comment:'上次登入Ip'"`
	Status            int       `gorm:"status" gorm:"comment:'状态'"`
	CreatedAt         time.Time `gorm:"status" gorm:"comment:'创建时间'"`
	UpdatedAt         time.Time `gorm:"status" gorm:"comment:'修改时间'"`
}
