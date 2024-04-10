package helpers

import (
	"time"
)

// Timer define métodos de la librería time
type Timer interface {
	FromString(date string) (time.Time, error)
	Now() *time.Time
}

// DefaultTimer permite acceder a métodos de Timer
type DefaultTimer struct{}

// FromString convierte una fecha string a un objeto time
func (d *DefaultTimer) FromString(date string) (time.Time, error) {
	return time.Parse(time.DateTime, date)
}

// Now retorna la fecha actual
func (d *DefaultTimer) Now() *time.Time {
	now := time.Now()
	return &now
}
