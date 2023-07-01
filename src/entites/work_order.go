package entites

import (
	"time"

	"github.com/google/uuid"
	"lucio.com/order-service/src/vo"
)

type WorkOrder struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID
	Title            string
	PlannedDateBegin time.Time
	PlannedDateEnd   time.Time
	Status           vo.WorkOrderStatus
	Type             vo.WorkOrderType
}
