package models

import (
	"time"

	"github.com/google/uuid"
	"lucio.com/order-service/src/vo"
)

type WorkOrder struct {
	ID               uuid.UUID          `gorm:"type:uuid;primaryKey"`
	CustomerID       uuid.UUID          `gorm:"type:uuid;not null"`
	Title            string             `gorm:"not null"`
	PlannedDateBegin time.Time          `gorm:"not null"`
	PlannedDateEnd   time.Time          `gorm:"not null"`
	Status           vo.WorkOrderStatus `gorm:"default:new"`
	CreatedAt        time.Time          `gorm:"default:now()"`
	Type             vo.WorkOrderType   `gorm:"default:orderService"`
}
