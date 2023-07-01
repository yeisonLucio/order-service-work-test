package vo

import (
	"errors"

	"lucio.com/order-service/src/helpers"
)

const (
	InactiveCustomerType string = "inactiveCustomer"
	OrderServiceType     string = "orderService"
)

type WorkOrderType struct {
	value string
}

var typeAllowed = []string{InactiveCustomerType, OrderServiceType}

func (w *WorkOrderType) SetValue(value string) error {
	if !helpers.StringContains(typeAllowed, value) {
		return errors.New("el tipo ingresado no es valido")
	}

	w.value = value

	return nil
}

func (w *WorkOrderType) GetValue() string {
	return w.value
}
