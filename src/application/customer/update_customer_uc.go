package customer

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/domain/customer/repositories"
)

// UpdateCustomerUC define las dependencias externas a utilizar en este caso de uso
type UpdateCustomerUC struct {
	CustomerRepository repositories.CustomerRepository
	Logger             *logrus.Logger
}

// Execute permite actualizar un customer
func (u *UpdateCustomerUC) Execute(
	updateCustomer entities.Customer,
) (*customerDtos.UpdatedCustomerResponse, *dtos.CustomError) {
	log := u.Logger.WithFields(logrus.Fields{
		"file":   "update_Customer_uc",
		"method": "Execute",
	})

	customer, err := u.CustomerRepository.FindByID(updateCustomer.ID)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
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

	log = log.WithField("customerToUpdate", customer)

	if err := u.CustomerRepository.Save(customer); err != nil {
		log.WithField("error", err)
		log.Error()
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
