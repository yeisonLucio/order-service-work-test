package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkOrder struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID       uuid.UUID `gorm:"type:uuid;not null"`
	Title            string    `gorm:"not null"`
	PlannedDateBegin time.Time `gorm:"not null"`
	PlannedDateEnd   time.Time `gorm:"not null"`
	Status           string    `gorm:"default:new"`
	CreatedAt        time.Time `gorm:"default:now()"`
	Type             string    `gorm:"default:orderService"`
}
