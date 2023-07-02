package dto

type WorkOrderDTO struct {
	ID               string
	Title            string      `json:"title"`
	PlannedDateBegin string      `json:"planned_date_begin"`
	PlannedDateEnd   string      `json:"planned_date_end"`
	WorkOrderType    string      `json:"work_order_type"`
	Status           string      `json:"status"`
	Customer         CustomerDTO `json:"customer" gorm:"embedded"`
}

type CustomerDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Status    bool   `json:"status"`
}
