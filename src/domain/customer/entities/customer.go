package entities

import "time"

type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Address   string
	StartDate *time.Time
	EndDate   *time.Time
	IsActive  bool
}
