package main

import (
	"service-api/internal/controllers"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/services", controllers.GetAllServices)
	route.GET("/service/:id", controllers.GetService)
	route.POST("/service", controllers.CreateService)
	route.PATCH("/services/:id", controllers.UpdateService)
	route.DELETE("/service/:id", controllers.DeleteService)

	route.GET("/user-service", controllers.GetUserServices)
	route.POST("/user-service", controllers.CreateUserService)
	route.DELETE("/user-service", controllers.DeleteService)

	route.Run()
}
