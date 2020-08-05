package services

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"gin-base/pkg/e"
	"gin-base/pkg/redis"
	"gin-base/pkg/utils"
)

const (
	vcodeRedisKey     = "gin-base:vcode:%s"
	vcodeDefaultValue = ""
	redisKeyTimeout   = 5 * time.Second
)

func GenVcode(phone string) (string, error) {
	vcode := utils.GenPhoneCode()

	key := fmt.Sprintf(vcodeRedisKey, phone)
	err := redis.Client.Set(redis.Ctx, key, vcode, redisKeyTimeout).Err()
	if err != nil {
		return vcodeDefaultValue, errors.Wrap(err, "redis set key error")
	}

	return vcode, nil
}

func CheckVcode(phone, vcode string) bool {
	oldVcode, err := GetVcode(phone)
	if err != nil {
		return false
	}

	if vcode != oldVcode {
		return false
	}

	return true
}

func GetVcode(phone string) (string, error) {
	key := fmt.Sprintf(vcodeRedisKey, phone)

	vcode, err := redis.Client.Get(redis.Ctx, key).Result()
	if err == redis.Nil {
		return vcodeDefaultValue, e.RedisKeyNotExist
	} else if err != nil {
		return vcodeDefaultValue, errors.Wrap(err, "redis get key error")
	}

	return vcode, nil
}
