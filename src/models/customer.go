package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName  string      `gorm:"not null" json:"first_name"`
	LastName   string      `gorm:"not null" json:"last_name"`
	Address    string      `gorm:"not null" json:"address"`
	StartDate  *time.Time  `json:"start_date"`
	EndDate    *time.Time  `json:"end_date"`
	IsActive   bool        `gorm:"default:false" json:"is_active"`
	CreateAt   time.Time   `gorm:"default:now()" json:"create_at"`
	WorkOrders []WorkOrder `gorm:"foreignKey:CustomerID" json:"-"`
}
