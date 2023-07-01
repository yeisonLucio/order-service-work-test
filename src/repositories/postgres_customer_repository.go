package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"lucio.com/order-service/src/entites"
	"lucio.com/order-service/src/models"
)

type PostgresCustomerRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresCustomerRepository) Save(customer entites.Customer) error {
	customerDB := models.Customer{
		ID:        uuid.MustParse(customer.ID),
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
	}

	result := p.ClientDB.Create(&customerDB)

	return result.Error
}
