package repositories

import (
	"errors"

	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/infra/models"
)

type PostgresCustomerRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresCustomerRepository) Create(customer *entities.Customer) *dtos.CustomError {
	var customerDB models.Customer
	customerDB.NewFromEntity(customer)

	result := p.ClientDB.Create(&customerDB)
	if result.Error != nil {
		return &dtos.CustomError{
			Code:  500,
			Error: result.Error,
		}
	}

	customer.ID = customerDB.ID.String()
	return nil
}

func (p *PostgresCustomerRepository) FindByID(ID string) (*entities.Customer, *dtos.CustomError) {
	var customer models.Customer

	if result := p.ClientDB.First(&customer, "id=?", ID); result.RowsAffected == 0 {
		return nil, &dtos.CustomError{
			Code:  404,
			Error: errors.New("cliente no encontrado"),
		}
	}

	return customer.ToEntity(), nil
}

func (p *PostgresCustomerRepository) Save(customer *entities.Customer) *dtos.CustomError {
	var customerDB models.Customer
	customerDB.FromEntity(customer)

	result := p.ClientDB.Updates(customerDB)
	if result.Error != nil {
		return &dtos.CustomError{
			Code:  500,
			Error: errors.New("error intentando actualizar"),
		}
	}

	return nil
}

func (p *PostgresCustomerRepository) GetActives() []entities.Customer {
	var customersDB []models.Customer

	p.ClientDB.
		Model(&models.Customer{}).
		Where("is_active", true).
		Scan(&customersDB)

	var customers []entities.Customer
	for _, item := range customersDB {
		customers = append(customers, *item.ToEntity())
	}

	return customers
}

func (p *PostgresCustomerRepository) DeleteByID(ID string) *dtos.CustomError {
	result := p.ClientDB.Delete(&models.Customer{}, "id=?", ID)
	if result.RowsAffected == 0 {
		return &dtos.CustomError{
			Code:  500,
			Error: errors.New("no fue posible eliminar el cliente"),
		}
	}

	return nil
}
