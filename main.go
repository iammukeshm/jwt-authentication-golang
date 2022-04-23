package main

import (
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:root@tcp(localhost:3306)/jwt_demo")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}
	}
	return router
}
