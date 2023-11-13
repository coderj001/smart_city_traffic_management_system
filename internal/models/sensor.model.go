package models

import (
	"time"

	"github.com/google/uuid"
)

// Sensor Table
type Sensor struct {
	ID           uuid.UUID    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Latitude     float64      `gorm:"type:float;not null" json:"latitude"`
	Longitude    float64      `gorm:"type:float;not null"  json:"longitude"`
	Rate         float64      `gorm:"type:float;not null"  json:"rate"`                                             // date output pre sec
	Status       bool         `gorm:"type:bool;not null;index:index_sensor_status,where:status=true" json:"status"` //TODO: need for fix check psql
	// Data         []SensorData `gorm:"foreignKey:SensorID;references:ID"`
	// IncidentData []Incident   `gorm:"foreignKey:SensorID;references:ID"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type CreateSensorRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Rate      float64 `json:"rate" binding:"required"`
	Status    bool    `json:"status,omitempty"`
}

type UpdateSensorRequest struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Rate      float64 `json:"rate,omitempty"`
	Status    bool    `json:"status,omitempty"`
}
