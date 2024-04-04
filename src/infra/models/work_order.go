package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/customer/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/domain/workorder/enums"
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
	Customer         Customer
}

func (w *WorkOrder) NewFromEntity(entity *entities.WorkOrder) {
	w.ID = uuid.New()
	w.CustomerID = uuid.MustParse(entity.CustomerID)
	w.Title = entity.Title
	w.PlannedDateBegin = entity.PlannedDateBegin
	w.PlannedDateEnd = entity.PlannedDateEnd
	w.Status = string(entity.Status)
	w.Type = string(entity.Type)
}

func (w *WorkOrder) ToEntity() *entities.WorkOrder {
	return &entities.WorkOrder{
		ID:               w.ID.String(),
		CustomerID:       w.CustomerID.String(),
		Title:            w.Title,
		Status:           enums.WorkOrderStatus(w.Status),
		Type:             enums.WorkOrderType(w.Type),
		PlannedDateBegin: w.PlannedDateBegin,
		PlannedDateEnd:   w.PlannedDateEnd,
	}
}

func (w *WorkOrder) FromEntity(entity *entities.WorkOrder) {
	w.ID = uuid.MustParse(entity.ID)
	w.CustomerID = uuid.MustParse(entity.CustomerID)
	w.Title = entity.Title
	w.PlannedDateBegin = entity.PlannedDateBegin
	w.PlannedDateEnd = entity.PlannedDateEnd
	w.Status = string(entity.Status)
	w.Type = string(entity.Type)
}

func (w *WorkOrder) ToWorkOrderDto() *workOrderDtos.WorkOrder {
	return &workOrderDtos.WorkOrder{
		ID:               w.ID.String(),
		Title:            w.Title,
		Type:             w.Type,
		Status:           w.Status,
		PlannedDateBegin: w.PlannedDateBegin.String(),
		PlannedDateEnd:   w.PlannedDateEnd.String(),
		Customer: dtos.Customer{
			ID:        w.Customer.ID.String(),
			FirstName: w.Customer.FirstName,
			LastName:  w.Customer.LastName,
			Address:   w.Customer.Address,
			StartDate: w.Customer.StartDate.String(),
			EndDate:   w.Customer.EndDate.String(),
			IsActive:  w.Customer.IsActive,
		},
	}
}
