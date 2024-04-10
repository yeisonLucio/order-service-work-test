package enums

type WorkOrderType string

const (
	//InactiveCustomer define el tipo inactive_customer de una orden de servicio
	InactiveCustomer WorkOrderType = "inactive_customer"

	//ServiceOrder define el tipo service_order de una orden de servicio
	ServiceOrder WorkOrderType = "service_order"
)
