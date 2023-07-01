package contracts

import "lucio.com/order-service/src/models"

type WorkOrderRepository interface {
	Create(workOrder models.WorkOrder) error
}
