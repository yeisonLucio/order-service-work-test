package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

type UpdateWorkOrderUC interface {
	Execute(updateWorkOrder entities.WorkOrder) (*workOrderDtos.UpdatedWorkOrderResponse, *dtos.CustomError)
}
