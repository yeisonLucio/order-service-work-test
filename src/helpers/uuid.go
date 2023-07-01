package helpers

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() string
}

type DefaultUUIDGenerator struct{}

func (g *DefaultUUIDGenerator) Generate() string {
	return uuid.New().String()
}
