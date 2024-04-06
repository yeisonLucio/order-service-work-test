package repositories

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/customer/entities"
	"lucio.com/order-service/src/infra/models"
)

type PostgresCustomerRepository struct {
	ClientDB *gorm.DB
	Logger   *logrus.Logger
}

func (p *PostgresCustomerRepository) Create(customer *entities.Customer) *dtos.CustomError {
	log := p.Logger.WithFields(logrus.Fields{
		"file":   "postgres_customer_repository",
		"method": "Create",
	})

	var customerDB models.Customer
	customerDB.NewFromEntity(customer)

	log = log.WithField("customer", customerDB)
	result := p.ClientDB.Create(&customerDB)
	if result.Error != nil {
		log.WithError(result.Error).Error()
		return &dtos.CustomError{
			Code:  500,
			Error: result.Error,
		}
	}

	customer.ID = customerDB.ID.String()
	return nil
}

func (p *PostgresCustomerRepository) FindByID(ID string) (*entities.Customer, *dtos.CustomError) {
	log := p.Logger.WithFields(logrus.Fields{
		"file":   "postgres_customer_repository",
		"method": "FindByID",
	})

	var customer models.Customer

	log = log.WithField("customer", customer)

	if result := p.ClientDB.First(&customer, "id=?", ID); result.RowsAffected == 0 {
		log.WithError(result.Error).Error()
		return nil, &dtos.CustomError{
			Code:  404,
			Error: errors.New("cliente no encontrado"),
		}
	}

	return customer.ToEntity(), nil
}

func (p *PostgresCustomerRepository) Save(customer *entities.Customer) *dtos.CustomError {
	log := p.Logger.WithFields(logrus.Fields{
		"file":   "postgres_customer_repository",
		"method": "Save",
	})

	var customerDB models.Customer
	customerDB.FromEntity(customer)

	log = log.WithField("customer", customerDB)

	result := p.ClientDB.Updates(customerDB)
	if result.Error != nil {
		log.WithError(result.Error).Error()
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
	log := p.Logger.WithFields(logrus.Fields{
		"file":   "postgres_customer_repository",
		"method": "DeleteByID",
	})

	result := p.ClientDB.Delete(&models.Customer{}, "id=?", ID)
	if result.RowsAffected == 0 {
		log.Error("no fue posible eliminar el cliente")
		return &dtos.CustomError{
			Code:  500,
			Error: errors.New("no fue posible eliminar el cliente"),
		}
	}

	return nil
}
