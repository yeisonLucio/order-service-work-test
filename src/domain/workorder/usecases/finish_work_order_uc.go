package contracts

import "lucio.com/order-service/src/domain/common/dtos"

// FinishWorkOrderUC define los métodos a implementar en el caso de uso de finalizar work order
type FinishWorkOrderUC interface {
	Execute(ID string) *dtos.CustomError
}
