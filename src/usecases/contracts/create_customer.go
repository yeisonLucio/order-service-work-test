package contracts

import (
	"lucio.com/order-service/src/dto"
)

type CreateCustomer interface {
	Execute(createCustomerDTO dto.CreateCustomerDTO) (*dto.CustomerCreatedResponse, error)
}
