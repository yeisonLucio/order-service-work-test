package contracts

import (
	"lucio.com/order-service/src/domain/common/dtos"
	customerDtos "lucio.com/order-service/src/domain/customer/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
)

// UpdateCustomerUC define los método a implementar en el caso de uso de actualizar un customer
type UpdateCustomerUC interface {
	Execute(updateCustomerDTO entities.Customer) (*customerDtos.UpdatedCustomerResponse, *dtos.CustomError)
}
