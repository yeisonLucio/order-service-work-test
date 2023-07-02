package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"lucio.com/order-service/src/helpers"
)

const (
	InactiveCustomerType string  = "inactive_customer"
	ServiceOrderType     string  = "service_order"
	StatusNew            string  = "new"
	StatusDone           string  = "done"
	StatusCancelled      string  = "cancelled"
	limitDifference      float64 = 2
)

var AllowedTypes = []string{InactiveCustomerType, ServiceOrderType}

type WorkOrder struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID       uuid.UUID  `gorm:"type:uuid;not null" json:"customer_id"`
	Title            string     `gorm:"not null" json:"title"`
	PlannedDateBegin *time.Time `gorm:"not null" json:"planned_date_begin"`
	PlannedDateEnd   *time.Time `gorm:"not null" json:"planned_date_end"`
	Status           string     `gorm:"default:new" json:"status"`
	CreatedAt        time.Time  `gorm:"default:now()" json:"create_at"`
	Type             string     `gorm:"default:orderService" json:"type"`
}

func (w *WorkOrder) Validate() error {
	if !helpers.StringContains(AllowedTypes, w.Type) {
		return errors.New("el tipo ingresado no esta permitido")
	}

	difference := w.PlannedDateEnd.Sub(*w.PlannedDateBegin)

	if difference.Hours() > limitDifference {
		return fmt.Errorf(
			"la diferencia de las fechas no puede ser mayor a %v horas",
			limitDifference,
		)
	}

	return nil
}
