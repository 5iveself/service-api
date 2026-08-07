package controllers

import (
	"net/http"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateUserValidate struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetAllUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(ctx *gin.Context) {
	var validate CreateUserValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Email: validate.Email, Username: validate.Username, Password: validate.Password}
	models.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
