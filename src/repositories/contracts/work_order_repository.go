package contracts

import (
	"lucio.com/order-service/src/dto"
	"lucio.com/order-service/src/models"
)

type WorkOrderRepository interface {
	Create(createWorkOrderDTO dto.CreateWorkOrderDTO) (*models.WorkOrder, error)
}
