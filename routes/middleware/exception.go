package middleware

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"

	"gin-base/app/api/common/helpers"
	"gin-base/pkg/e"
	. "gin-base/pkg/global"
)

func Exception() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				message := "----------------------------------\n"
				message += fmt.Sprintf("*RequestId:* %s\n", c.MustGet("RequestId").(string))
				message += fmt.Sprintf("*Link:* %s%s\n", c.Request.Host, c.Request.URL)
				message += fmt.Sprintf("*Project:* %s\n", App.Conf.Project.Name)
				message += fmt.Sprintf("*Environment:* %s\n", App.Conf.Project.RunMode)
				message += fmt.Sprintf("*RequestIp:* %s\n", c.ClientIP())
				message += fmt.Sprintf("*RequestMethod:* %s\n", c.Request.Method)
				message += fmt.Sprintf("*Time:* %s\n", time.Now().Format("2006-01-02 15:04:05"))
				message += fmt.Sprintf("*Exception:* %s\n", err)
				message += fmt.Sprintf("*DebugStack:* %s=\n", string(debug.Stack()))

				log.Println(message)

				helpers.SendResponse(c, e.InternalServerError, nil)
			}
		}()

		c.Next()
	}
}
