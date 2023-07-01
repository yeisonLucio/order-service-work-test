package usecases

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/models"
	"lucio.com/order-service/src/repositories/contracts"
)

type CreateCustomerUC struct {
	CustomerRepository contracts.CustomerRepository
	UUID               helpers.UUIDGenerator
}

func (c *CreateCustomerUC) Execute(
	createCustomerDTO dto.CreateCustomerDTO,
) (*dto.CreatedCustomerDTO, error) {
	customer := models.Customer{
		ID:        c.UUID.Generate(),
		FirstName: createCustomerDTO.FirstName,
		LastName:  createCustomerDTO.LastName,
		Address:   createCustomerDTO.Address,
	}

	if err := c.CustomerRepository.Create(customer); err != nil {
		return nil, err
	}

	response := dto.CreatedCustomerDTO{
		ID:        customer.ID.String(),
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
	}

	return &response, nil
}
