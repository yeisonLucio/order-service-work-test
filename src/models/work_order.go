package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"lucio.com/order-service/src/helpers"
)

const (
	InactiveCustomerType string  = "inactive_customer"
	ServiceOrderType     string  = "service_order"
	StatusNew            string  = "new"
	StatusDone           string  = "done"
	StatusCancelled      string  = "cancelled"
	dateRange            float64 = 2
)

var (
	ErrInvalidType     = errors.New("el tipo ingresado no esta permitido")
	ErrInvalidDates    = fmt.Errorf("rango de fechas invalido, no puede ser mayor que %v horas", dateRange)
	ErrBeginDateFormat = errors.New("formato incorrecto para planned_date_begin")
	ErrEndDateFormat   = errors.New("formato incorrecto para planned_date_end")
	AllowedTypes       = []string{InactiveCustomerType, ServiceOrderType}
)

type WorkOrder struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID       uuid.UUID  `gorm:"type:uuid;not null" json:"customer_id"`
	Title            string     `gorm:"not null" json:"title"`
	PlannedDateBegin *time.Time `gorm:"not null" json:"planned_date_begin"`
	PlannedDateEnd   *time.Time `gorm:"not null" json:"planned_date_end"`
	Status           string     `gorm:"default:new" json:"status"`
	CreatedAt        time.Time  `gorm:"default:now()" json:"create_at"`
	Type             string     `gorm:"default:orderService" json:"type"`
	DeletedAt        gorm.DeletedAt
}

func (w *WorkOrder) Validate() error {
	if !helpers.StringContains(AllowedTypes, w.Type) {
		return ErrInvalidType
	}

	difference := w.PlannedDateEnd.Sub(*w.PlannedDateBegin)

	if difference.Hours() > dateRange {
		return ErrInvalidDates
	}

	return nil
}

func (w *WorkOrder) SetPlannedDateBegin(date string) error {
	beginPlannedDate, err := time.Parse(time.DateTime, date)
	if err != nil {
		return ErrBeginDateFormat
	}
	w.PlannedDateBegin = &beginPlannedDate

	return nil
}

func (w *WorkOrder) SetPlannedDateEnd(date string) error {
	endPlannedDate, err := time.Parse(time.DateTime, date)
	if err != nil {
		return ErrEndDateFormat
	}
	w.PlannedDateEnd = &endPlannedDate

	return nil
}
