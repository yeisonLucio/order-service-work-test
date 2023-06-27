package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type PostgresCustomerRepository struct {
	ClientDB *gorm.DB
}

func (p PostgresCustomerRepository) Create(createCustomerDTO dto.CreateCustomerDTO) (*models.Customer, error) {

	customer := models.Customer{
		FirstName: createCustomerDTO.FirstName,
		LastName:  createCustomerDTO.LastName,
		Address:   createCustomerDTO.Address,
	}

	if result := p.ClientDB.Create(&customer); result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}
