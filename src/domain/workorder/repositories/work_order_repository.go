package repositories

import (
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

type WorkOrderRepository interface {
	Create(workOrder *entities.WorkOrder) *dtos.CustomError
	IsTheFirstOrder(customerID string) bool
	FindByID(ID string) (*entities.WorkOrder, *dtos.CustomError)
	Save(workOrder *entities.WorkOrder) *dtos.CustomError
	GetFiltered(filters workOrderDtos.WorkOrderFilters) []*workOrderDtos.WorkOrder
	DeleteByID(ID string) *dtos.CustomError
}
