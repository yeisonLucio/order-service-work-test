package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
)

type UpdateCustomerUC interface {
	Execute(updateCustomerDTO entities.Customer) (*customerDtos.UpdatedCustomerResponse, *dtos.CustomError)
}
