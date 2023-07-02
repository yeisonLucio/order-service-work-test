package contracts

import "lucio.com/order-service/src/dto"

type UpdateWorkOrderUC interface {
	Execute(updateWorkOrder dto.UpdateWorkOrder) (*dto.UpdatedWorkOrder, error)
}
