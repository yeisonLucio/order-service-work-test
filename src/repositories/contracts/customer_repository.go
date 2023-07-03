package contracts

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type CustomerRepository interface {
	Create(customer models.Customer) error
	FindByID(ID string) (*models.Customer, error)
	GetByID(ID string) (*dto.CustomerDTO, error)
	Save(customer *models.Customer) error
	GetActives() []dto.CustomerDTO
	DeleteByID(ID string) error
}
