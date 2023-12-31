package routes

import (
	"backend/handlers"
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.UserLoginWithPassword)
	r.GET("/",  handlers.Home)
	r.GET("/categories", handlers.Categories)
	r.GET("/quizes", middlewares.UserAuthMiddleware(), handlers.Quizes)
	r.POST("/score", middlewares.UserAuthMiddleware(), handlers.ScoreTracking)
}
