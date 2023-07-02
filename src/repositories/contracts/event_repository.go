package contracts

import "lucio.com/order-service/src/models"

type EventRepository interface {
	NotifyWorkOrderFinished(payload models.WorkOrder) error
}
