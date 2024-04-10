package dto

// ErrorResponse define estructura básica de respuesta de error
type ErrorResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
} // @name errorResponse
