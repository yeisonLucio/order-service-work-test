package contracts

import (
	"lucio.com/order-service/src/dto"
)

type CreateCustomerUC interface {
	Execute(createCustomerDTO dto.CreateCustomerDTO) (*dto.CreatedCustomerDTO, error)
}
