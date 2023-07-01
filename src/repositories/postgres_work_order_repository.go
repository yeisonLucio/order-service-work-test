package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type PostgresCreateWorkOrderRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresCreateWorkOrderRepository) Create(
	createWorkOrderDTO dto.CreateWorkOrderDTO,
) (*models.WorkOrder, error) {

	beginDate, err := time.Parse(time.RFC3339, createWorkOrderDTO.PlannedDateBegin)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse(time.RFC3339, createWorkOrderDTO.PlannedDateEnd)
	if err != nil {
		return nil, err
	}

	workOrder := models.WorkOrder{
		CustomerID:       uuid.MustParse(createWorkOrderDTO.CustomerID),
		Title:            createWorkOrderDTO.Title,
		PlannedDateBegin: beginDate,
		PlannedDateEnd:   endDate,
		Type:             models.WorkOrderType(createWorkOrderDTO.WorkOrderType),
	}

	if result := p.ClientDB.Create(&workOrder); result.Error != nil {
		return nil, result.Error
	}

	return &workOrder, nil
}
