package dtos

type CreatedCustomerResponse struct {
	ID        string `json:"id"`
	IsActive  bool   `json:"is_active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}
