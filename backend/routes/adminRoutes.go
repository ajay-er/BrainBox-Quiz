package routes

import (
	"backend/handlers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	r.POST("/login",middlewares.AdminAuthMiddleware(),handlers.AdminLogin)
}
