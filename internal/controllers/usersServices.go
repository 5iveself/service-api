package controllers

import (
	"net/http"
	"service-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserServiceValidate struct {
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	ServiceID uint      `json:"service_id" binding:"required"`
	StartDate string    `json:"start_date" binding:"required"`
	ExpDate   string    `json:"exp_date" binding:"required"`
}

type UserServiceResponse struct {
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
}

func GetUserServices(ctx *gin.Context) {
	var userServices []models.UserService

	query := models.DB.Preload("Service")

	if id := ctx.Query("id"); id != "" {
		query = query.Where("user_id = ?", id)
	}

	if name := ctx.Query("name"); name != "" {
		query = query.Joins("JOIN services ON services.id = user_services.service_id").
			Where("services.service_name ILIKE ?", "%"+name+"%")
	}

	if err := query.Find(&userServices).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]UserServiceResponse, 0, len(userServices))
	for _, us := range userServices {
		response = append(response, UserServiceResponse{
			ServiceName: us.Service.ServiceName,
			Price:       us.Service.Price,
			UserID:      us.UserID,
			StartDate:   us.StartDate,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"service": response})
}

func CreateUserService(ctx *gin.Context) {
	var validate CreateUserServiceValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := models.UserService{ServiceID: validate.ServiceID, UserID: validate.UserID, StartDate: validate.StartDate, ExpDate: validate.ExpDate}
	models.DB.Create(&userService)

	ctx.JSON(http.StatusOK, gin.H{"user-service": userService})
}

func DeleteUserService(ctx *gin.Context) {
	var userService models.UserService

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&userService).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	models.DB.Delete(&userService)
	ctx.JSON(http.StatusOK, gin.H{"service": true})
}
