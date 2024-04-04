package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"lucio.com/order-service/src/domain/customer/entities"
)

type Customer struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey"`
	FirstName  string      `gorm:"not null"`
	LastName   string      `gorm:"not null"`
	Address    string      `gorm:"not null"`
	StartDate  *time.Time  `json:"start_date"`
	EndDate    *time.Time  `json:"end_date"`
	IsActive   bool        `gorm:"default:false"`
	CreateAt   time.Time   `gorm:"default:now()"`
	WorkOrders []WorkOrder `gorm:"foreignKey:CustomerID"`
	DeletedAt  gorm.DeletedAt
}

func (c *Customer) ToEntity() *entities.Customer {
	return &entities.Customer{
		ID:        c.ID.String(),
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Address:   c.LastName,
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		IsActive:  c.IsActive,
	}
}

func (c *Customer) NewFromEntity(entity *entities.Customer) {
	c.ID = uuid.New()
	c.FirstName = entity.FirstName
	c.LastName = entity.LastName
	c.Address = entity.Address
}

func (c *Customer) FromEntity(entity *entities.Customer) {
	c.ID = uuid.MustParse(entity.ID)
	c.FirstName = entity.FirstName
	c.LastName = entity.LastName
	c.Address = entity.Address
}
