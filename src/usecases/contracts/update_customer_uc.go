package contracts

import (
	"lucio.com/order-service/src/dto"
)

type UpdateCustomerUC interface {
	Execute(updateCustomerDTO dto.UpdateCustomerDTO) (*dto.CustomerDTO, error)
}
