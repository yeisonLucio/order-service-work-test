package dtos

type CustomError struct {
	Code  int
	Error error
	Data  interface{}
}
