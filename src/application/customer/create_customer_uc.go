package customer

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/domain/customer/repositories"
)

type CreateCustomerUC struct {
	CustomerRepository repositories.CustomerRepository
	Logger             *logrus.Logger
}

func (c *CreateCustomerUC) Execute(
	createCustomer entities.Customer,
) (*customerDtos.CreatedCustomerResponse, *dtos.CustomError) {
	log := c.Logger.WithFields(logrus.Fields{
		"file":   "create_Customer_uc",
		"method": "Execute",
	})

	err := c.CustomerRepository.Create(&createCustomer)
	if err != nil {
		log = log.WithField("error", err)
		log.Error()
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
