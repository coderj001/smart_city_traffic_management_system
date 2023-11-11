package controllers

import (
	"net/http"

	"github.com/coderj001/smart_city_traffic_management_system/internal/models"
	"github.com/coderj001/smart_city_traffic_management_system/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This will consider as traffic data

type IncidentController struct {
	DB *gorm.DB
}

func NewIncidentController(db *gorm.DB) *IncidentController {
	return &IncidentController{DB: db}
}

func (i *IncidentController) GetIncident(ctx *gin.Context) {
	incidentID := ctx.Param("id")
	var incident models.Incident
	if err := i.DB.Where("id = ?", incidentID).First(&incident).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get incident",
		"data":    incident,
	})
}

func (i *IncidentController) GetIncidents(ctx *gin.Context) {
	var allIncidents []models.Incident

	query := i.DB
	statusParams := ctx.Query("status")
	if statusParams != "" {
		query = query.Where("status = ?", statusParams)
	}
	typeParams := ctx.Query("type")
	if typeParams != "" {
		query = query.Where("type = ?", typeParams)
	}

	if err := query.Find(&allIncidents).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all incidents",
		"data":    allIncidents,
	})
}

func (i *IncidentController) CreateIncident(ctx *gin.Context) {
	var payload models.IncidentDataPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	newIncident := models.Incident{
		SensorID:    payload.SensorID,
		Type:        payload.Type,
		Description: payload.Description,
		Status:      payload.Status,
	}

	if err := i.DB.Create(&newIncident).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Incident is reported successfully!"})
}

func (i *IncidentController) UpdateIncident(ctx *gin.Context) {
	incidentID := ctx.Param("id")
	var incident models.Incident
	var newIncident models.IncidentDataUpdate

	if err := ctx.ShouldBindJSON(&newIncident); err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := i.DB.Where("id = ?", incidentID).First(&incident).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	updatedIncident := models.Incident{
		Type:        newIncident.Type,
		Description: newIncident.Description,
		Status:      newIncident.Status,
	}

	if err := i.DB.Model(&incident).Updates(updatedIncident).Error; err != nil {
		services.HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Incident is updated successfully!",
		"data":    updatedIncident,
	})
}
