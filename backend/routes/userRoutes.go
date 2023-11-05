package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.Signup)

}
