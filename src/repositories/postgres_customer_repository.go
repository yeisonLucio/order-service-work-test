package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"lucio.com/order-service/src/models"
)

type PostgresCustomerRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresCustomerRepository) Create(customer models.Customer) error {
	result := p.ClientDB.Create(&customer)

	return result.Error
}

func (p *PostgresCustomerRepository) FindByID(ID string) *models.Customer {
	customer := models.Customer{
		ID: uuid.MustParse(ID),
	}

	if result := p.ClientDB.First(&customer); result.RowsAffected == 0 {
		return nil
	}

	return &customer
}

func (p *PostgresCustomerRepository) Save(customer *models.Customer) error {
	result := p.ClientDB.Save(&customer)
	return result.Error
}
