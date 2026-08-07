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

type ServiceResponse struct {
	ID          uint   `json:"id"`
	ServiceName string `json:"title"`
	Price       int    `json:"price"`
}

// GetAllServices godoc
// @Summary Get all services
// @Description Returns a list of all services
// @Tags services
// @Produce json
// @Success 200 {object} map[string][]ServiceResponse
// @Router /services [get]
func GetAllServices(ctx *gin.Context) {
	var service []models.Service
	models.DB.Find(&service)

	ctx.JSON(http.StatusOK, gin.H{"services": service})
}

// GetService godoc
// @Summary Get a service by ID
// @Description Returns a single service by its ID
// @Tags services
// @Produce json
// @Param id path uint true "Service ID"
// @Success 200 {object} map[string]ServiceResponse
// @Failure 400 {object} map[string]string
// @Router /service/{id} [get]
func GetService(ctx *gin.Context) {
	var user models.Service

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Service not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// CreateService godoc
// @Summary Create a new service
// @Description Creates a new service and returns it
// @Tags services
// @Accept json
// @Produce json
// @Param service body CreateServiceValidate true "Service data"
// @Success 200 {object} map[string]ServiceResponse
// @Failure 400 {object} map[string]string
// @Router /service [post]
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

// UpdateService godoc
// @Summary Update a service
// @Description Updates an existing service by its ID
// @Tags services
// @Accept json
// @Produce json
// @Param id path uint true "Service ID"
// @Param service body UpdateServiceValidate true "Service data"
// @Success 200 {object} map[string]ServiceResponse
// @Failure 400 {object} map[string]string
// @Router /services/{id} [patch]
func UpdateService(ctx *gin.Context) {
	var service models.Service

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&service).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Service not found"})
		return
	}

	var validate UpdateServiceValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&service).Updates(validate)

	ctx.JSON(http.StatusOK, gin.H{"service": service})
}

// DeleteService godoc
// @Summary Delete a service
// @Description Deletes a service by its ID
// @Tags services
// @Produce json
// @Param id path uint true "Service ID"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Router /service/{id} [delete]
func DeleteService(ctx *gin.Context) {
	var service models.Service

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&service).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Service not found"})
		return
	}
	models.DB.Delete(&service)
	ctx.JSON(http.StatusOK, gin.H{"service": true})
}
