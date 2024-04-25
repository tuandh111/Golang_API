package main

import (
	"gocrudapibackend/controllers"
	"gocrudapibackend/initializers"
	"gocrudapibackend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvVariable()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/api/v1/signup", controllers.Signup)
	r.POST("/api/v1/login", controllers.Login)
	r.GET("/api/v1/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
