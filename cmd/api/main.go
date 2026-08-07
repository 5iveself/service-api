package main

import (
	"log/slog"
	"os"

	_ "service-api/docs"

	"service-api/internal/controllers"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Service API
// @version 1.0
// @description API for managing services and user subscriptions.
// @host localhost:8080
// @BasePath /
func main() {

	//Handling error for making dir logs
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic(err)
	}

	//open file for logging
	file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	//guarantee file closing
	defer file.Close()

	logger := slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(logger)

	gin.DefaultWriter = file
	route := gin.Default()

	//Connect my db
	models.ConnectDB()

	slog.Info("Database connected")

	//Swagger route
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Services routes
	route.GET("/services", controllers.GetAllServices)
	route.GET("/service/:id", controllers.GetService)
	route.POST("/service", controllers.CreateService)
	route.PATCH("/services/:id", controllers.UpdateService)
	route.DELETE("/service/:id", controllers.DeleteService)

	//Subscriptions routes
	route.GET("/user-service", controllers.GetUserServices)
	route.POST("/user-service", controllers.CreateUserService)
	route.DELETE("/user-service/:id", controllers.DeleteUserService)

	//Logging about server status
	slog.Info("Server starting", "addr", ":8080")
	if err := route.Run(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
