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
		Status:    payload.Status,
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

func (sc *SensorController) GetAll(ctx *gin.Context) {
	var allSensor []models.Sensor

	query := sc.DB

	// if status is true, then filter all sensor with status true
	if statusParams, exists := ctx.GetQuery("status"); exists {
		var status bool = false
		if statusParams == "true" {
			status = true
		}

		// Apply filter
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&allSensor).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
    return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Traffic Sensor",
		"data":    allSensor,
	})
}

func (sc *SensorController) GetByID(ctx *gin.Context) {
	sensorID := ctx.Param("id")
	var sensor models.Sensor

	if err := sc.DB.First(&sensor, "id = ?", sensorID).Error; err != nil {
		services.HandleError(ctx, http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sensor",
		"data":    sensor,
	})
}

func (sc *SensorController) UpdateByID(ctx *gin.Context) {
	sensorID := ctx.Param("id")
	var sensor models.Sensor
	var newSensor models.UpdateSensorRequest

	if err := ctx.ShouldBindJSON(&newSensor); err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
	}

	if err := sc.DB.First(&sensor, "id = ?", sensorID).Error; err != nil {
		services.HandleError(ctx, http.StatusNotFound, err)
	}

	updateSensor := models.Sensor{
		Latitude:  newSensor.Latitude,
		Longitude: newSensor.Longitude,
		Rate:      newSensor.Rate,
		Status:    newSensor.Status,
	}

	if err := sc.DB.Model(&sensor).Updates(updateSensor).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
	}

	ctx.JSON(
    http.StatusOK,
    gin.H{
      "message": "Sensor Updated Successfully",
      "data":    sensor,
    },
  )

}
