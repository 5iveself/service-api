package controllers

import (
	"net/http"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateServiceValidate struct {
	ServiceName string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

type UpdateServiceValidate struct {
	ServiceName string `json:"title"`
	Price       int    `json:"price"`
}

func GetAllServices(ctx *gin.Context) {
	var service []models.Service
	models.DB.Find(&service)

	ctx.JSON(http.StatusOK, gin.H{"services": service})
}

func GetService(ctx *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Service not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateService(ctx *gin.Context) {
	var validate CreateServiceValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := models.Service{ServiceName: validate.ServiceName, Price: validate.Price}
	models.DB.Create(&service)

	ctx.JSON(http.StatusOK, gin.H{"service": service})
}

func UpdateService(ctx *gin.Context) {
	var service models.Service

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&service).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Service not found"})
		return
	}

	var validate UpdateUserValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&service).Updates(validate)

	ctx.JSON(http.StatusOK, gin.H{"service": service})
}

func DeleteService(ctx *gin.Context) {
	var service models.Service

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&service).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	models.DB.Delete(&service)
	ctx.JSON(http.StatusOK, gin.H{"service": true})
}
