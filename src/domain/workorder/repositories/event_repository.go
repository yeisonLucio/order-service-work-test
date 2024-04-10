package repositories

import (
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

// EventRepository define los métodos a implementar por el repositorio de eventos
type EventRepository interface {
	NotifyWorkOrderFinished(payload *entities.WorkOrder) *dtos.CustomError
}
