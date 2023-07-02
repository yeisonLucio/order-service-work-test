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
	query := p.ClientDB.Table("work_orders wo").
		Select(`wo.id,
			wo.title,
			wo.planned_date_begin,
			wo.planned_date_end,
			wo.status, 
			wo.type,
			wo.customer_id,
			c.id,
			c.first_name,
			c.last_name, 
			c.address,
			c.start_date,
			c.end_date,
			c.is_active`).
		Joins("JOIN customers c ON c.id = wo.customer_id")

	if filters.PlannedDateBegin != "" && filters.PlannedDateEnd != "" {
		query = query.Where("planned_date_begin > ?", filters.PlannedDateBegin).
			Where("planned_date_end < ?", filters.PlannedDateEnd)
	}

	if filters.Status != "" {
		query = query.Where("status", filters.Status)
	}

	query.Scan(&workOrders)

	return workOrders
}

func (p *PostgresWorkOrderRepository) GetByCustomerID(customerID string) []dto.WorkOrderDTO {
	var workOrders []dto.WorkOrderDTO

	p.ClientDB.Table("work_orders wo").
		Select(`wo.id,
			wo.title,
			wo.planned_date_begin,
			wo.planned_date_end,
			wo.status, 
			wo.type,
			wo.customer_id,
			c.id,
			c.first_name,
			c.last_name, 
			c.address,
			c.start_date,
			c.end_date,
			c.is_active`).
		Joins("JOIN customers c ON c.id = wo.customer_id").
		Where("wo.deleted_at IS NULL").
		Where("wo.customer_id", customerID).
		Scan(&workOrders)

	return workOrders
}

func (p *PostgresWorkOrderRepository) DeleteByID(ID string) error {
	result := p.ClientDB.Delete(&models.WorkOrder{}, "id=?", ID)
	if result.RowsAffected == 0 {
		return errors.New("el id ingresado no existe")
	}

	return nil
}

func (p *PostgresWorkOrderRepository) FindByIdWithCustomer(ID string) (*dto.WorkOrderDTO, error) {
	var workOrder dto.WorkOrderDTO

	result := p.ClientDB.Table("work_orders wo").
		Select(`wo.id,
			wo.title,
			wo.planned_date_begin,
			wo.planned_date_end,
			wo.status, 
			wo.type,
			wo.customer_id,
			c.id,
			c.first_name,
			c.last_name, 
			c.address,
			c.start_date,
			c.end_date,
			c.is_active`).
		Joins("JOIN customers c ON c.id = wo.customer_id").
		Where("wo.deleted_at IS NULL").
		Where("wo.id", ID).
		Scan(&workOrder)

	if result.RowsAffected == 0 {
		return nil, errors.New("orden de servicio no encontrada")
	}

	return &workOrder, nil
}
