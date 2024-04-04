package customer

import (
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/domain/customer/repositories"
)

type UpdateCustomerUC struct {
	CustomerRepository repositories.CustomerRepository
}

func (u *UpdateCustomerUC) Execute(
	updateCustomer entities.Customer,
) (*customerDtos.UpdatedCustomerResponse, *dtos.CustomError) {
	customer, err := u.CustomerRepository.FindByID(updateCustomer.ID)
	if err != nil {
		return nil, err
	}

	if updateCustomer.Address != "" {
		customer.Address = updateCustomer.Address
	}

	if updateCustomer.FirstName != "" {
		customer.FirstName = updateCustomer.FirstName
	}

	if updateCustomer.LastName != "" {
		customer.LastName = updateCustomer.LastName
	}

	if err := u.CustomerRepository.Save(customer); err != nil {
		return nil, err
	}

	var startDate, endDate string

	if customer.StartDate != nil {
		startDate = customer.StartDate.String()
	}

	if customer.EndDate != nil {
		endDate = customer.EndDate.String()
	}

	return &customerDtos.UpdatedCustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
		StartDate: startDate,
		EndDate:   endDate,
		IsActive:  customer.IsActive,
	}, nil
}
