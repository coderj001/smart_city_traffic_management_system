package models

import (
	"time"

	"github.com/google/uuid"
)

// Traffic Data model
type SensorData struct {
	Id           uint      `gorm:"primaryKey"`
	SensorID     uuid.UUID `gorm:"type:uuid;not null;index"` // This is the foreign key
	VericleCount int
	AverageSpeed float64
	TimeStamp    time.Time
}
