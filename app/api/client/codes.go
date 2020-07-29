package client

import (
	"context"
	"fmt"
	"gin-base/app/api/common/helpers"
	"github.com/gin-gonic/gin"
	"time"

	"gin-base/pkg/e"
	"gin-base/pkg/redis"
	"gin-base/pkg/utils"
)

var ctx = context.Background()

// 获取手机验证码
// parameter
//   phone
func GetPhoneVerifyCode(c *gin.Context) {
	phone := c.Query("phone")
	if phone == "" {
		helpers.SendResponse(c, e.ParamPhoneEmpty, nil)

		return
	}

	// 验证手机号

	key := fmt.Sprintf("phone:%s", phone)

	_, err := redis.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// key does not exist
	} else if err != nil {
		// error
	} else {
		// key exist
		helpers.SendResponse(c, e.ParamPhoneEmpty, nil)

		return
	}

	// 生成验证码
	code := utils.GenPhoneCode()
	fmt.Println(code)

	redis.Client.Set(ctx, key, code, 60*time.Second)

	// 发送短信
	helpers.SendResponse(c, nil, nil)
}
