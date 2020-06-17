package client

import (
	"gin-base/app/api/client/helpers"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	helpers.SendResponse(c, nil, nil)
}
