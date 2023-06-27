package usecases

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/repositories/contracts"
)

type CreateCustomerUC struct {
	CustomerRepository contracts.CustomerRepository
}

func (c *CreateCustomerUC) Execute(
	createCustomerDTO dto.CreateCustomerDTO,
) (*dto.CustomerCreatedResponse, error) {

	customerDB, err := c.CustomerRepository.Create(createCustomerDTO)
	if err != nil {
		return nil, err
	}

	response := dto.CustomerCreatedResponse{
		ID:        customerDB.ID.String(),
		FirstName: customerDB.FirstName,
		LastName:  customerDB.LastName,
		Address:   customerDB.Address,
	}

	return &response, nil
}
