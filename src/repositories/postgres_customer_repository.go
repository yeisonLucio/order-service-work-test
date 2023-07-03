package repositories

import (
	"errors"

	"gorm.io/gorm"
	"lucio.com/order-service/src/dto"
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

func (p *PostgresCustomerRepository) GetActives() []dto.CustomerDTO {
	var customers []dto.CustomerDTO
	p.ClientDB.Model(&models.Customer{}).Where("is_active", true).Scan(&customers)
	return customers
}

func (p *PostgresCustomerRepository) DeleteByID(ID string) error {
	result := p.ClientDB.Delete(&models.Customer{}, "id=?", ID)
	if result.RowsAffected == 0 {
		return errors.New("el id ingresado no existe")
	}

	return nil
}

func (p *PostgresCustomerRepository) GetByID(ID string) (*dto.CustomerDTO, error) {
	var customer dto.CustomerDTO

	result := p.ClientDB.
		Model(&models.Customer{}).
		Where("id", ID).
		Scan(&customer)

	if result.RowsAffected == 0 {
		return nil, errors.New("cliente no encontrado")
	}

	return &customer, nil
}
