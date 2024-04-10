package dtos

import "lucio.com/order-service/src/domain/customer/dtos"

// WorkOrder define objeto de transferencia de una orden de servicio
type WorkOrder struct {
	ID               string        `json:"id"`
	Title            string        `json:"title"`
	PlannedDateBegin string        `json:"planned_date_begin"`
	PlannedDateEnd   string        `json:"planned_date_end"`
	Type             string        `json:"type"`
	Status           string        `json:"status"`
	Customer         dtos.Customer `json:"customer"`
}
