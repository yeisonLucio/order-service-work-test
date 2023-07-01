package vo

import (
	"errors"

	"lucio.com/order-service/src/helpers"
)

const (
	StatusNew       = "new"
	StatusDone      = "done"
	StatusCancelled = "cancelled"
)

type WorkOrderStatus struct {
	value string
}

var statusAllowed = []string{StatusNew, StatusDone, StatusCancelled}

func (w *WorkOrderStatus) SetValue(status string) error {
	if !helpers.StringContains(statusAllowed, status) {
		return errors.New("el status ingresado no es valido")
	}

	w.value = status

	return nil
}

func (w *WorkOrderStatus) GetValue() string {
	return w.value
}
