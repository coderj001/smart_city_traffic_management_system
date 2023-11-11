package controllers

import (
	"net/http"
	"time"

	"github.com/coderj001/smart_city_traffic_management_system/internal/models"
	"github.com/coderj001/smart_city_traffic_management_system/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This will consider as traffic data

type SensorDataController struct {
	DB *gorm.DB
}

func NewSensorDataController(db *gorm.DB) *SensorDataController {
	return &SensorDataController{DB: db}
}

func (c *SensorDataController) GetSensorData(ctx *gin.Context) {
	sensorDataID := ctx.Param("id")
	var sensorData models.SensorData

	if err := c.DB.First(&sensorData, "id = ?", sensorDataID).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Sensor data fetched successfully",
		"data":    sensorData,
	})
}

func (c *SensorDataController) PostSensorData(ctx *gin.Context) {
  var payload models.SensorDataPayload
  if err := ctx.ShouldBindJSON(&payload); err != nil {
    services.HandleError(ctx, http.StatusBadRequest, err)
    return
  }

  newSensorData := models.SensorData{
    SensorID: payload.SensorID,
    VericleCount: payload.VericleCount,
    AverageSpeed: payload.AverageSpeed,
    TimeStamp: time.Now(),
  }

  if err := c.DB.Create(&newSensorData).Error; err != nil {
    services.HandleError(ctx, http.StatusBadRequest, err)
    return
  }

	ctx.JSON(http.StatusCreated,
		gin.H{"message": "Sensor data posted successfully"},
	)
}
