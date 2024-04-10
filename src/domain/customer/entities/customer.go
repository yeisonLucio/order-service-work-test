package entities

import "time"

// Customer define la entidad de un customer
type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Address   string
	StartDate *time.Time
	EndDate   *time.Time
	IsActive  bool
}
