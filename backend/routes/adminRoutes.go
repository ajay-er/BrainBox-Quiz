package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	r.POST("/login", handlers.AdminLogin)
	r.GET("/dashboard",handlers.AdminDashboard)
}
