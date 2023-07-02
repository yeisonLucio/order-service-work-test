package contracts

import (
	"lucio.com/order-service/src/models"
)

type CustomerRepository interface {
	Create(customer models.Customer) error
	FindByID(ID string) (*models.Customer, error)
	Save(customer *models.Customer) error
	GetActives() []models.Customer
}
