package models

import (
	"time"

	"github.com/google/uuid"
)

// Traffic Data model
type SensorData struct {
	Id           uint      `gorm:"primaryKey"`
	SensorID     uuid.UUID `gorm:"type:uuid;not null;index;"` // This is the foreign key
	Sensor       Sensor    `gorm:"foreignKey:SensorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VehicleCount int
	AverageSpeed float64
	TimeStamp    time.Time `gorm:"autoCreateTime;autoUpdateTime"`
}

type SensorDataPayload struct {
	SensorID     uuid.UUID `json:"sensor_id" binding:"required"`
	VehicleCount int       `json:"vehicle_count,omitempty"`
	AverageSpeed float64   `json:"average_speed,omitempty"`
}
