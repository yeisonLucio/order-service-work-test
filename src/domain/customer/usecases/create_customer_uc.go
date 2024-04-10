package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
)

// CreateCustomerUC define los método a implementar en el caso de uso de crear un customer
type CreateCustomerUC interface {
	Execute(customer entities.Customer) (*customerDtos.CreatedCustomerResponse, *dtos.CustomError)
}
