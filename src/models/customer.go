package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrInvalidCustomerFirstName = errors.New("el campo first_name es requerido")
	ErrInvalidCustomerLastName  = errors.New("el campo last_name es requerido")
	ErrInvalidCustomerAddress   = errors.New("el campo address es requerido")
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
	DeletedAt  gorm.DeletedAt
}

func (c *Customer) Validate() error {
	if c.FirstName == "" {
		return ErrInvalidCustomerFirstName
	}
	if c.LastName == "" {
		return ErrInvalidCustomerLastName
	}
	if c.Address == "" {
		return ErrInvalidCustomerAddress
	}

	return nil
}
