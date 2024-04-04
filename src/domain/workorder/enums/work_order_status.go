package enums

type WorkOrderStatus string

const (
	StatusNew       WorkOrderStatus = "new"
	StatusDone      WorkOrderStatus = "done"
	StatusCancelled WorkOrderStatus = "cancelled"
)
