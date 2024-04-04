package repositories

import (
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

type EventRepository interface {
	NotifyWorkOrderFinished(payload *entities.WorkOrder) *dtos.CustomError
}
