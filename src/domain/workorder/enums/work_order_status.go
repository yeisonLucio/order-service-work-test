package enums

type WorkOrderStatus string

const (
	//StatusNew define estado new de una orden de servicio
	StatusNew WorkOrderStatus = "new"

	//StatusDone define estado done de una orden de servicio
	StatusDone WorkOrderStatus = "done"

	//StatusCancelled define estado cancelled de una orden de servicio
	StatusCancelled WorkOrderStatus = "cancelled"
)
