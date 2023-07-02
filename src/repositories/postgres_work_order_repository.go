package repositories

import (
	"errors"

	"gorm.io/gorm"
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type PostgresWorkOrderRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresWorkOrderRepository) Create(workOrder models.WorkOrder) error {
	result := p.ClientDB.Create(&workOrder)

	return result.Error
}

func (p *PostgresWorkOrderRepository) IsTheFirstOrder(WorkOrderID, customerID string) bool {
	var result string
	p.ClientDB.Model(models.WorkOrder{}).
		Select("id").
		Where("customer_id", customerID).
		Order("created_at").
		Limit(1).
		Scan(&result)

	return result == WorkOrderID
}

func (p *PostgresWorkOrderRepository) FindByID(ID string) (*models.WorkOrder, error) {
	var workOrder models.WorkOrder

	if result := p.ClientDB.First(&workOrder, "id=?", ID); result.RowsAffected == 0 {
		return nil, errors.New("orden de servicio no encontrada")
	}

	return &workOrder, nil
}

func (p *PostgresWorkOrderRepository) Save(workOrder *models.WorkOrder) error {
	result := p.ClientDB.Updates(workOrder)
	return result.Error
}

func (p *PostgresWorkOrderRepository) GetAll(filters dto.WorkOrderFilters) []dto.WorkOrderDTO {
	var workOrders []dto.WorkOrderDTO
	query := p.ClientDB.Table("work_orders").
		Select("work_orders.*, customers.*")

	if filters.PlannedDateBegin != "" && filters.PlannedDateEnd != "" {
		query = query.Where("planned_date_begin > ?", filters.PlannedDateBegin).
			Where("planned_date_end < ?", filters.PlannedDateEnd)
	}

	if filters.Status != "" {
		query = query.Where("status", filters.Status)
	}

	query.Joins("JOIN customers ON customers.id = work_orders.customer_id").Scan(&workOrders)
	return workOrders
}
