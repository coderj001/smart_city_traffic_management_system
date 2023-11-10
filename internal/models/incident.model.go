package models

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	ID          uint      `gorm:"primaryKey"`
	SensorID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Type        string    `gorm:"not null"`
	Description string    `gorm:"varchar(255)"`
	Status      string    `gorm:"not null;default:open"`
	TimeStamp   time.Time
}
