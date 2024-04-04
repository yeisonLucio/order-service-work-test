package factories

import "lucio.com/order-service/src/domain/common/helpers"

func NewTimeLib() helpers.Timer {
	return &helpers.DefaultTimer{}
}
