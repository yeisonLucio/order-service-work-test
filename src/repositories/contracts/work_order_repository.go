package contracts

import (
	"lucio.com/order-service/src/entites"
)

type WorkOrderRepository interface {
	Save(workOrder entites.WorkOrder) error
}
