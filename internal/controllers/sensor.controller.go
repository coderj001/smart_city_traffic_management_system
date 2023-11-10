package controllers

import (
	"net/http"

	"github.com/coderj001/smart_city_traffic_management_system/internal/models"
	"github.com/coderj001/smart_city_traffic_management_system/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SensorController struct {
	DB *gorm.DB
}

func NewSensorController(db *gorm.DB) *SensorController {
	return &SensorController{DB: db}
}

func (sc *SensorController) Create(ctx *gin.Context) {
	var payload models.CreateSensorRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	newSensor := models.Sensor{
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		Rate:      payload.Rate,
	}
	if err := sc.DB.Create(&newSensor).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
	}
	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "New Traffice Sensor Register Successfully! ✅",
			"data":    newSensor,
		},
	)
}
