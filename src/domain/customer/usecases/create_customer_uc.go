package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
)

type CreateCustomerUC interface {
	Execute(customer entities.Customer) (*customerDtos.CreatedCustomerResponse, *dtos.CustomError)
}
