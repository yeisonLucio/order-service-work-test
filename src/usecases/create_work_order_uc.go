package usecases

import "lucio.com/order-service/src/dto"

type CreateWorkOrderUC struct {
}

func (c *CreateWorkOrderUC) Execute() (*dto.OrderCreatedResponseDTO, error) {
	return nil, nil
}
