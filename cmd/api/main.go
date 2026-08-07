package main

import (
	"net/http"
	"service-api/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	database.ConnectDB()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Test!"})
	})

	route.Run()
}
