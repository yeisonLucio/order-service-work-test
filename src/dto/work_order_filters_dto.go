package dto

type WorkOrderFilters struct {
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Status           string `json:"status"`
}
