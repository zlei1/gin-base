package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"gin-base/pkg/e"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := e.DecodeE(err)

	c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
