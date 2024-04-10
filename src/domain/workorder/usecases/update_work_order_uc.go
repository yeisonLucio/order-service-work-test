package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

// UpdateWorkOrderUC define los métodos a implementar en el caso de uso de actualizar una orden de servicio
type UpdateWorkOrderUC interface {
	Execute(updateWorkOrder entities.WorkOrder) (*workOrderDtos.UpdatedWorkOrderResponse, *dtos.CustomError)
}
