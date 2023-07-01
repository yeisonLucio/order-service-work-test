package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/models"
)

type PostgresCreateWorkOrderRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresCreateWorkOrderRepository) Create(workOrder models.WorkOrder) error {
	result := p.ClientDB.Create(&workOrder)

	return result.Error
}
