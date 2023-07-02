package dto

type WorkOrderDTO struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	PlannedDateBegin string      `json:"planned_date_begin"`
	PlannedDateEnd   string      `json:"planned_date_end"`
	Type             string      `json:"type"`
	Status           string      `json:"status"`
	Customer         CustomerDTO `json:"customer" gorm:"embedded"`
}

type CreateWorkOrderDTO struct {
	CustomerID       string `json:"customer_id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type"`
}

type CreatedWorkOrderDTO struct {
	ID               string `json:"id"`
	Status           string `json:"status"`
	CustomerID       string `json:"customer_id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type"`
}

type WorkOrderFilters struct {
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Status           string `json:"status"`
}

type UpdateWorkOrder struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type"`
}

type UpdatedWorkOrder struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	PlannedDateBegin string `json:"planned_date_begin"`
	PlannedDateEnd   string `json:"planned_date_end"`
	Type             string `json:"type"`
	Status           string `json:"status"`
}
