package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-base/app/api/common/helpers"
	"gin-base/pkg/e"
	"gin-base/pkg/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth_header := c.Request.Header.Get("Authorization")

		if len(auth_header) == 0 {
			helpers.SendResponse(c, e.TokenInvalid, nil)

			c.Abort()
			return
		}

		jctx, err := jwt.ParseToken(auth_header)

		if err != nil {
			helpers.SendResponse(c, err, nil)
			c.Abort()
			return
		}

		c.Set("user_id", jctx.UserID)

		c.Next()
	}
}
