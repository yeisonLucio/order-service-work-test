package repositories

import (
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
)

// CustomerRepository define los métodos a implementar por el repositorio de customer
type CustomerRepository interface {
	Create(customer *entities.Customer) *dtos.CustomError
	FindByID(ID string) (*entities.Customer, *dtos.CustomError)
	Save(customer *entities.Customer) *dtos.CustomError
	GetActives() []entities.Customer
	DeleteByID(ID string) *dtos.CustomError
}
