package routes

import (
	"github.com/gin-gonic/gin"

	admin_api "gin-base/app/api/admin"
	"gin-base/routes/middleware"
)

func InitAdmin(router *gin.Engine) {
	r := router.Group("/api/admin")
	r.POST("/sessions", admin_api.Login)

	auth := r.Group("")
	auth.Use(middleware.Auth())

	auth.GET("/admins", admin_api.IndexAdmin)
	auth.GET("/admins/:id", admin_api.ShowAdmin)
	auth.POST("/admins", admin_api.CreateAdmin)
	auth.PUT("/admins/:id", admin_api.UpdateAdmin)
	auth.DELETE("/admins/:id", admin_api.DeleteAdmin)
}
