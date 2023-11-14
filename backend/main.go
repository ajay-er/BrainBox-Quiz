package main

import (
	"backend/config"
	database "backend/db"
	"backend/routes"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading the config file")
		fmt.Println("hey")
	}
	DB, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Database is ", DB)
	router := gin.Default()

	corss := cors.DefaultConfig()
	corss.AllowOrigins = []string{"*"}
	corss.AllowMethods = []string{"GET", "POST","PUT","PATCH","DELETE"}
	router.Use(cors.New(corss))
	routes.UserRoutes(router.Group("/users"))
	routes.AdminRoutes(router.Group("/admin"))
	router.Run(cfg.BASE_URL)

}
