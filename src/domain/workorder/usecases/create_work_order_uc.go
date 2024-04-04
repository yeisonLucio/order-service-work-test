package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

type CreateWorkOrderUC interface {
	Execute(createWorkOrderDTO entities.WorkOrder) (*workOrderDtos.CreatedWorkOrderResponse, *dtos.CustomError)
}
