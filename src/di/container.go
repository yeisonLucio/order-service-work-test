package di

import (
	"lucio.com/order-service/src/controllers"
	"lucio.com/order-service/src/database"
	"lucio.com/order-service/src/helpers"
	"lucio.com/order-service/src/repositories"
	"lucio.com/order-service/src/usecases"
)

type Dependencies struct {
	CustomerController  *controllers.CustomerController
	WorkOrderController *controllers.WorkOrderController
}

var Container Dependencies

func BuildContainer() {
	// libraries

	time := &helpers.DefaultTimer{}
	uuid := &helpers.DefaultUUIDGenerator{}

	// repositories

	customerRepository := &repositories.PostgresCustomerRepository{
		ClientDB: database.DB,
	}

	workOrderRepository := &repositories.PostgresWorkOrderRepository{
		ClientDB: database.DB,
	}

	// use cases

	createCustomerUC := &usecases.CreateCustomerUC{
		CustomerRepository: customerRepository,
		UUID:               uuid,
	}

	createWorkOrderUC := &usecases.CreateWorkOrderUC{
		WorkOrderRepository: workOrderRepository,
		CustomerRepository:  customerRepository,
		UUID:                uuid,
		Time:                time,
	}

	// controllers

	Container.CustomerController = &controllers.CustomerController{
		CreateCustomerUC:  createCustomerUC,
		CreateWorkOrderUC: createWorkOrderUC,
	}

	Container.WorkOrderController = &controllers.WorkOrderController{}
}
