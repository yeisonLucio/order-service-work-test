package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	Address    string    `gorm:"not null"`
	StartDate  *time.Time
	EndDate    *time.Time
	IsActive   bool        `gorm:"default:false"`
	CreateAt   time.Time   `gorm:"default:now()"`
	WorkOrders []WorkOrder `gorm:"foreignKey:CustomerID"`
}
