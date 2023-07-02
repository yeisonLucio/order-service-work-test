package contracts

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type WorkOrderRepository interface {
	Create(workOrder models.WorkOrder) error
	IsTheFirstOrder(WorkOrderID, customerID string) bool
	FindByID(ID string) (*models.WorkOrder, error)
	Save(workOrder *models.WorkOrder) error
	GetAll(filters dto.WorkOrderFilters) []dto.WorkOrderDTO
}
