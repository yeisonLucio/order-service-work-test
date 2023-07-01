package contracts

import (
	"lucio.com/order-service/src/entites"
)

type CustomerRepository interface {
	Save(customer entites.Customer) error
}
