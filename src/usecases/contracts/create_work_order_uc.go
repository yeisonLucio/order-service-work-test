package contracts

import "lucio.com/order-service/src/dto"

type CreateWorkOrderUC interface {
	Execute(createWorkOrderDTO dto.CreateWorkOrderDTO) (*dto.CreatedWorkOrderDTO, error)
}
