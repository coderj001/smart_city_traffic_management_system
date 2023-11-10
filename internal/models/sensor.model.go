package models

import (
	"time"

	"github.com/google/uuid"
)

type Sensor struct {
	ID           uuid.UUID    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Latitude     float64      `gorm:"type:float;not null"`
	Longitude    float64      `gorm:"type:float;not null"`
	Rate         float64      `gorm:"type:float;not null"`
	Status       bool         `gorm:"type:bool;not null;default:true;index:index_status,where:status=true"`
	Data         []SensorData `gorm:"foreignKey:SensorID;references:ID"`
	IncidentData []Incident   `gorm:"foreignKey:SensorID;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
