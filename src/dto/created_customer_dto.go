package dto

type CreatedCustomerDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Status    bool   `json:"status"`
}
