package dto

type FinishWorkOrderDTO struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	WorkOrderType    string `json:"work_order_type"`
}