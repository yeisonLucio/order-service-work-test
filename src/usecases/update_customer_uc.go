package usecases

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/repositories/contracts"
)

type UpdateCustomerUC struct {
	CustomerRepository contracts.CustomerRepository
}

func (u *UpdateCustomerUC) Execute(
	updateCustomerDTO dto.UpdateCustomerDTO,
) (*dto.CustomerDTO, error) {

	customer, err := u.CustomerRepository.FindByID(updateCustomerDTO.ID)
	if err != nil {
		return nil, err
	}

	customer.Address = updateCustomerDTO.Address
	customer.FirstName = updateCustomerDTO.FirstName
	customer.LastName = updateCustomerDTO.LastName

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

	return &dto.CustomerDTO{
		ID:        customer.ID.String(),
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
		StartDate: startDate,
		EndDate:   endDate,
		IsActive:  customer.IsActive,
	}, nil
}
