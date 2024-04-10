package entities

import (
	"time"

	"lucio.com/order-service/src/domain/workorder/enums"
)

// WorkOrder define la entidad de una orden de servicio
type WorkOrder struct {
	ID               string
	CustomerID       string
	Title            string
	Status           enums.WorkOrderStatus
	Type             enums.WorkOrderType
	PlannedDateBegin *time.Time
	PlannedDateEnd   *time.Time
}
