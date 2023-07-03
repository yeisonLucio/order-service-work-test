package dto

type ErrorResponse struct {
	ID      string `json:"id" enums:"bad_request,unexpected_error,record_not_found" example:"string"`
	Message string `json:"message"`
} // @name errorResponse
