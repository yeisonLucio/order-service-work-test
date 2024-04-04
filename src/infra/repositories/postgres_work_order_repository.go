package repositories

import (
	"errors"

	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/common/dtos"
	workOrderDtos "lucio.com/order-service/src/domain/workorder/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
	"lucio.com/order-service/src/infra/models"
)

type PostgresWorkOrderRepository struct {
	ClientDB *gorm.DB
}

func (p *PostgresWorkOrderRepository) Create(workOrder *entities.WorkOrder) *dtos.CustomError {
	var workOrderDB models.WorkOrder
	workOrderDB.NewFromEntity(workOrder)

	result := p.ClientDB.Create(&workOrderDB)
	if result.Error != nil {
		return &dtos.CustomError{
			Code:  500,
			Error: result.Error,
		}
	}

	workOrder.ID = workOrderDB.ID.String()

	return nil
}

func (p *PostgresWorkOrderRepository) IsTheFirstOrder(customerID string) bool {
	var totalOrders int64
	p.ClientDB.Model(models.WorkOrder{}).
		Where("customer_id", customerID).
		Count(&totalOrders)

	return totalOrders == 1
}

func (p *PostgresWorkOrderRepository) FindByID(ID string) (*entities.WorkOrder, *dtos.CustomError) {
	var workOrder models.WorkOrder

	if result := p.ClientDB.First(&workOrder, "id=?", ID); result.RowsAffected == 0 {
		return nil, &dtos.CustomError{
			Code:  404,
			Error: errors.New("orden de servicio no encontrada"),
		}
	}

	return workOrder.ToEntity(), nil
}

func (p *PostgresWorkOrderRepository) Save(workOrder *entities.WorkOrder) *dtos.CustomError {
	var workOrderDB models.WorkOrder
	workOrderDB.FromEntity(workOrder)

	result := p.ClientDB.Updates(workOrder)
	if result.Error != nil {
		return &dtos.CustomError{
			Code:  500,
			Error: result.Error,
		}
	}

	if result.RowsAffected == 0 {
		return &dtos.CustomError{
			Code:  500,
			Error: errors.New("orden de servicio no actualizada"),
		}
	}

	return nil
}

func (p *PostgresWorkOrderRepository) GetFiltered(filters workOrderDtos.WorkOrderFilters) []*workOrderDtos.WorkOrder {
	var workOrders []models.WorkOrder
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
		Joins("JOIN customers c ON c.id = wo.customer_id").
		Where("wo.deleted_at IS NULL")

	if filters.PlannedDateBegin != "" && filters.PlannedDateEnd != "" {
		query = query.Where("planned_date_begin > ?", filters.PlannedDateBegin).
			Where("planned_date_end < ?", filters.PlannedDateEnd)
	}

	if filters.Status != "" {
		query = query.Where("status", filters.Status)
	}

	if filters.CustomerID != "" {
		query = query.Where("wo.customer_id", filters.CustomerID)
	}

	if filters.ID != "" {
		query = query.Where("wo.id", filters.ID)
	}

	query.Scan(&workOrders)

	var workOrderDto []*workOrderDtos.WorkOrder
	for _, v := range workOrders {
		workOrderDto = append(workOrderDto, v.ToWorkOrderDto())
	}

	return workOrderDto
}

func (p *PostgresWorkOrderRepository) DeleteByID(ID string) *dtos.CustomError {
	result := p.ClientDB.Delete(&models.WorkOrder{}, "id=?", ID)
	if result.RowsAffected == 0 {
		return &dtos.CustomError{
			Code:  404,
			Error: errors.New("el id ingresado no existe"),
		}
	}

	return nil
}
