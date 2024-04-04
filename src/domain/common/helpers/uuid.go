package helpers

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() uuid.UUID
}

type DefaultUUIDGenerator struct{}

func (g *DefaultUUIDGenerator) Generate() uuid.UUID {
	return uuid.New()
}
