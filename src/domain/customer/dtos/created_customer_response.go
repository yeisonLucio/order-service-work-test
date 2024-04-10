package dtos

// CreatedCustomerResponse define el modelo de respuesta para cuando se crea un customer
type CreatedCustomerResponse struct {
	ID        string `json:"id"`
	IsActive  bool   `json:"is_active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}
