package helpers

import "github.com/google/uuid"

// UUIDGenerator define métodos a implementar de la librería uuid
type UUIDGenerator interface {
	Generate() uuid.UUID
}

// DefaultUUIDGenerator permite utilizar métodos para generar uuid
type DefaultUUIDGenerator struct{}

// Generate genera un nuevo uuid
func (g *DefaultUUIDGenerator) Generate() uuid.UUID {
	return uuid.New()
}
