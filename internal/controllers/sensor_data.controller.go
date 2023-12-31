package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/coderj001/smart_city_traffic_management_system/internal/models"
	"github.com/coderj001/smart_city_traffic_management_system/internal/services"
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
	go processSensorData(payload, c.DB)

	ctx.JSON(http.StatusCreated,
		gin.H{"message": "Sensor data posted successfully"},
	)
}

func processSensorData(data models.SensorDataPayload, db *gorm.DB) {

	newSensorData := models.SensorData{
		SensorID:     data.SensorID,
		VehicleCount: data.VehicleCount,
		AverageSpeed: data.AverageSpeed,
		TimeStamp:    time.Now(),
	}

	if err := db.Create(&newSensorData).Error; err != nil {
		panic(err)
	}
}
