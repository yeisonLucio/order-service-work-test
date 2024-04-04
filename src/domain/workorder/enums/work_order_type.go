package enums

type WorkOrderType string

const (
	InactiveCustomer WorkOrderType = "inactive_customer"
	ServiceOrder     WorkOrderType = "service_order"
)
