package middleware

import (
	"gin-base/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("RequestId")

		if requestId == "" {
			requestId = utils.GenUuid()
		}

		c.Set("RequestId", requestId)
		c.Writer.Header().Set("RequestId", requestId)

		c.Next()
	}
}
