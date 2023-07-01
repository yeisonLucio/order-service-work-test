package models

import (
	"time"

	"github.com/google/uuid"
)

type WorkOrderStatus string
type WorkOrderType string

const (
	StatusNew            WorkOrderStatus = "new"
	StatusDone           WorkOrderStatus = "done"
	StatusCancelled      WorkOrderStatus = "cancelled"
	InactiveCustomerType WorkOrderType   = "inactiveCustomer"
	OrderServiceType     WorkOrderType   = "orderService"
)

type WorkOrder struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID       uuid.UUID       `gorm:"not null"`
	Title            string          `gorm:"not null"`
	PlannedDateBegin time.Time       `gorm:"not null"`
	PlannedDateEnd   time.Time       `gorm:"not null"`
	Status           WorkOrderStatus `gorm:"default:new"`
	CreatedAt        time.Time       `gorm:"default:now()"`
	Type             WorkOrderType   `gorm:"default:orderService"`
}
