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

type UpdateUserValidate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAllUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(ctx *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
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

func UpdateUser(ctx *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var validate UpdateUserValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(validate)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
