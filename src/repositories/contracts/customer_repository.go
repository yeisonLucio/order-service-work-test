package contracts

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type CustomerRepository interface {
	Create(createCustomerDTO dto.CreateCustomerDTO) (*models.Customer, error)
}
