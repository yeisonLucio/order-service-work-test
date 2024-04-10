package dtos

// CustomError estructura de error
type CustomError struct {
	Code  int
	Error error
	Data  interface{}
}
