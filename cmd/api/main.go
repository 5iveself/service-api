package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Test!"})
	})

	route.Run()
}
