package customer

import (
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/domain/customer/repositories"
)

type CreateCustomerUC struct {
	CustomerRepository repositories.CustomerRepository
}

func (c *CreateCustomerUC) Execute(
	createCustomer entities.Customer,
) (*customerDtos.CreatedCustomerResponse, *dtos.CustomError) {
	err := c.CustomerRepository.Create(&createCustomer)
	if err != nil {
		return nil, err
	}

	response := customerDtos.CreatedCustomerResponse{
		ID:        createCustomer.ID,
		FirstName: createCustomer.FirstName,
		LastName:  createCustomer.LastName,
		Address:   createCustomer.Address,
	}

	return &response, nil
}
