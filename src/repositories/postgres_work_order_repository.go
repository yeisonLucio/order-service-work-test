package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/models"
)

type PostgresWorkOrderRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresWorkOrderRepository) Create(workOrder models.WorkOrder) error {
	result := p.ClientDB.Create(&workOrder)

	return result.Error
}
