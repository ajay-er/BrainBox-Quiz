package routes

import (
	"backend/handlers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {

	r.POST("/login", handlers.AdminLogin)
	r.Use(middlewares.AuthorizationMiddleware())
	{

		r.GET("/dashboard", handlers.AdminDashboard)
		users := r.Group("users")
		{
			users.GET("/:page", handlers.GetUsers)
			users.POST("", handlers.Signup)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
			users.GET("/block-users/:id", handlers.BlockUser)
			users.GET("/un-block-users/:id", handlers.UnBlockUser)

		}
		category := r.Group("/category")
		{
			category.POST("", handlers.AddCategory)
			category.PUT("", handlers.UpdateCategory)
			category.DELETE("/:id", handlers.DeleteCategory)
		}
	}

}
