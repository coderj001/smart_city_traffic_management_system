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
	VericleCount int
	AverageSpeed float64
	TimeStamp    time.Time
}

type SensorDataPayload struct {
	SensorID     uuid.UUID  `json:"sensor_id" binding:"required"`
	VericleCount int     `json:"vericle_count,omitempty"`
	AverageSpeed float64 `json:"average_speed,omitempty"`
}
