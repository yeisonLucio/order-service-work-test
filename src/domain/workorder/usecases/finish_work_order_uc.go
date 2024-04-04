package contracts

import "lucio.com/order-service/src/domain/common/dtos"

type FinishWorkOrderUC interface {
	Execute(ID string) *dtos.CustomError
}
