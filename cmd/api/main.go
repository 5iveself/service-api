package main

import (
	"service-api/internal/controllers"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/users", controllers.GetAllUsers)
	route.GET("/user/:id", controllers.GetUser)
	route.POST("/user", controllers.CreateUser)
	route.PATCH("/users/:id", controllers.UpdateUser)

	route.Run()
}
