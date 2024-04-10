package dtos

// UpdatedCustomerResponse define modelo de respuesta de cuando se actualiza un customer
type UpdatedCustomerResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	IsActive  bool   `json:"is_active"`
}
