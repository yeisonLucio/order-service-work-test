package models

import (
	"time"

	"github.com/google/uuid"
)

type workOrderStatus string
type orderType string

const (
	StatusNew            workOrderStatus = "new"
	StatusDone           workOrderStatus = "done"
	StatusCancelled      workOrderStatus = "cancelled"
	InactiveCustomerType orderType       = "inactiveCustomer"
	OrderServiceType     orderType       = "orderService"
)

type WorkOrder struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID       uuid.UUID       `gorm:"not null"`
	Title            string          `gorm:"not null"`
	PlannedDateBegin time.Time       `gorm:"not null"`
	PlannedDateEnd   time.Time       `gorm:"not null"`
	Status           workOrderStatus `gorm:"default:new"`
	CreatedAt        time.Time       `gorm:"default:now()"`
	Type             orderType       `gorm:"default:orderService"`
}
