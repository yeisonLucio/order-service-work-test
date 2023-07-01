package usecases

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/entites"
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/repositories/contracts"
)

type CreateCustomerUC struct {
	CustomerRepository contracts.CustomerRepository
	UUID               helpers.DefaultUUIDGenerator
}

func (c *CreateCustomerUC) Execute(
	createCustomerDTO dto.CreateCustomerDTO,
) (*dto.CustomerCreatedResponse, error) {
	customer := entites.Customer{
		ID:        c.UUID.Generate(),
		FirstName: createCustomerDTO.FirstName,
		LastName:  createCustomerDTO.LastName,
		Address:   createCustomerDTO.Address,
	}

	if err := c.CustomerRepository.Save(customer); err != nil {
		return nil, err
	}

	response := dto.CustomerCreatedResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
	}

	return &response, nil
}
