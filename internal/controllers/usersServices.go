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

// GetUserServices godoc
// @Summary Get user services
// @Description Returns the services subscribed by a user, optionally filtered by user id or service name
// @Tags user-services
// @Produce json
// @Param id query string false "Filter by user id"
// @Param name query string false "Filter by service name (partial match)"
// @Success 200 {object} map[string][]UserServiceResponse
// @Failure 500 {object} map[string]string
// @Router /user-service [get]
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

// CreateUserService godoc
// @Summary Subscribe a user to a service
// @Description Creates a new user service subscription
// @Tags user-services
// @Accept json
// @Produce json
// @Param userService body CreateUserServiceValidate true "Subscription data"
// @Success 200 {object} map[string]models.UserService
// @Failure 400 {object} map[string]string
// @Router /user-service [post]
func CreateUserService(ctx *gin.Context) {
	var validate CreateUserServiceValidate
	if err := ctx.ShouldBindJSON(&validate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := models.UserService{ServiceID: validate.ServiceID, UserID: validate.UserID, StartDate: validate.StartDate, ExpDate: validate.ExpDate}
	models.DB.Create(&userService)

	ctx.JSON(http.StatusOK, gin.H{"user-service": "create-succeess!"})
}

// DeleteUserService godoc
// @Summary Delete a user service
// @Description Deletes a user service subscription by its ID
// @Tags user-services
// @Produce json
// @Param id path uint true "Subscription ID"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Router /user-service/{id} [delete]
func DeleteUserService(ctx *gin.Context) {
	var userService models.UserService

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&userService).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user-service not found"})
		return
	}
	models.DB.Delete(&userService)
	ctx.JSON(http.StatusOK, gin.H{"service": true})
}
