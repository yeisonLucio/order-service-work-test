package dtos

type CreatedWorkOrderResponse struct {
	ID               string `json:"id"`
	Status           string `json:"status"`
	CustomerID       string `json:"customer_id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type"`
}
