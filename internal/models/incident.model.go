package models

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	ID          uint      `gorm:"primaryKey"`
	SensorID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Sensor      Sensor    `gorm:"foreignKey:SensorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Type        string    `gorm:"not null"` // accidents, roadblocks
	Description string    `gorm:"varchar(255)"`
	Status      string    `gorm:"not null"` // resolved, ongoing
	TimeStamp   time.Time
}

type IncidentDataPayload struct {
	SensorID    uuid.UUID `json:"sensor_id" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status" binding:"required"`
}
