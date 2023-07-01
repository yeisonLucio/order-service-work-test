package repositories

import (
	"gorm.io/gorm"
	"lucio.com/order-service/src/entites"
	"lucio.com/order-service/src/models"
)

type PostgresCreateWorkOrderRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresCreateWorkOrderRepository) Save(workOrder entites.WorkOrder) error {
	workOrderDB := models.WorkOrder{
		ID:               workOrder.ID,
		CustomerID:       workOrder.CustomerID,
		Title:            workOrder.Title,
		PlannedDateBegin: workOrder.PlannedDateBegin,
		PlannedDateEnd:   workOrder.PlannedDateEnd,
		Type:             workOrder.Title,
	}

	result := p.ClientDB.Create(&workOrderDB)

	return result.Error
}
