package services

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"gin-base/pkg/global"
	"gin-base/pkg/utils"
)

type PhoneVcodeService interface {
	GenVcode(phone string) (string, error)
	CheckVcode(phone, vcode string) bool
	GetVcode(phone string) (string, error)
}

var PhoneVcodeSvc = NewPhoneVcodeService()

type phoneVcodeService struct{}

func NewPhoneVcodeService() PhoneVcodeService {
	return &phoneVcodeService{}
}

const (
	vcodeRedisKey     = "gin-base:vcode:%s"
	vcodeDefaultValue = ""
	redisKeyTimeout   = 60 * time.Second
)

func (server *phoneVcodeService) GenVcode(phone string) (string, error) {
	vcode := utils.GenPhoneCode()

	key := fmt.Sprintf(vcodeRedisKey, phone)
	err := global.App.RedisClient.Set(context.Background(), key, vcode, redisKeyTimeout).Err()
	if err != nil {
		return vcodeDefaultValue, errors.Wrap(err, "redis set key error")
	}

	return vcode, nil
}

func (server *phoneVcodeService) CheckVcode(phone, vcode string) bool {
	oldVcode, err := server.GetVcode(phone)
	if err != nil {
		return false
	}

	if vcode != oldVcode {
		return false
	}

	return true
}

func (server *phoneVcodeService) GetVcode(phone string) (string, error) {
	key := fmt.Sprintf(vcodeRedisKey, phone)

	vcode, err := global.App.RedisClient.Get(context.Background(), key).Result()
	if err == global.RedisClientNil {
		return "", nil
	} else if err != nil {
		return "", errors.Wrap(err, "redis get key error")
	}

	return vcode, nil
}
