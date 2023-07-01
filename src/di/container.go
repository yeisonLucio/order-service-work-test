package di

import (
	"lucio.com/order-service/src/controllers"
	"lucio.com/order-service/src/database"
	"lucio.com/order-service/src/repositories"
	"lucio.com/order-service/src/usecases"
)

type Dependencies struct {
	CustomerController  *controllers.CustomerController
	WorkOrderController *controllers.WorkOrderController
}

var Container Dependencies

func BuildContainer() {
	// repositories

	customerRepository := &repositories.PostgresCustomerRepository{
		ClientDB: database.DB,
	}

	// use cases

	createCustomerUC := &usecases.CreateCustomerUC{
		CustomerRepository: customerRepository,
	}

	// controllers

	Container.CustomerController = &controllers.CustomerController{
		CreateCustomerUC: createCustomerUC,
	}
}
