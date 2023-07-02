package repositories

import (
	"errors"

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

func (p *PostgresCustomerRepository) FindByID(ID string) (*models.Customer, error) {
	var customer models.Customer

	if result := p.ClientDB.First(&customer, "id=?", ID); result.RowsAffected == 0 {
		return nil, errors.New("cliente no encontrado")
	}

	return &customer, nil
}

func (p *PostgresCustomerRepository) Save(customer *models.Customer) error {
	result := p.ClientDB.Updates(customer)
	return result.Error
}

func (p *PostgresCustomerRepository) GetActives() []models.Customer {
	var customers []models.Customer
	p.ClientDB.Where("is_active", true).Find(&customers)
	return customers
}
