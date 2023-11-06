package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.UserLoginWithPassword)
	r.GET("/", handlers.Home)

}
