package dtos

type UpdatedWorkOrderResponse struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type"`
	Status           string `json:"status"`
}
