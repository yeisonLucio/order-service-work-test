package helpers

import (
	"time"
)

type Timer interface {
	FromString(date string) (time.Time, error)
	Now() *time.Time
}

type DefaultTimer struct{}

func (d *DefaultTimer) FromString(date string) (time.Time, error) {
	return time.Parse(time.DateTime, date)
}

func (d *DefaultTimer) Now() *time.Time {
	now := time.Now()
	return &now
}
