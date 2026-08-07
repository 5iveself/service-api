package controllers

import (
	"net/http"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
