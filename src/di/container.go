package di

import (
	"lucio.com/order-service/src/config/postgres"
	"lucio.com/order-service/src/config/redis"
	"lucio.com/order-service/src/controllers"
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
		ClientDB: postgres.DB,
	}

	workOrderRepository := &repositories.PostgresWorkOrderRepository{
		ClientDB: postgres.DB,
	}

	eventRepository := &repositories.RedisEventRepository{
		RedisClient: redis.RedisClient,
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

	finishWorkOrderUC := &usecases.FinishWorkOrderUC{
		WorkOrderRepository: workOrderRepository,
		CustomerRepository:  customerRepository,
		EventRepository:     eventRepository,
		Time:                time,
	}

	// controllers

	Container.CustomerController = &controllers.CustomerController{
		CreateCustomerUC:   createCustomerUC,
		CreateWorkOrderUC:  createWorkOrderUC,
		CustomerRepository: customerRepository,
	}

	Container.WorkOrderController = &controllers.WorkOrderController{
		FinishWorkOrderUC:   finishWorkOrderUC,
		WorkOrderRepository: workOrderRepository,
	}
}
