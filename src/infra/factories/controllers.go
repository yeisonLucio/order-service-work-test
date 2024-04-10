package factories

import "lucio.com/order-service/src/infra/controllers"

// NewCustomerController función para inicializar el controlador de los clientes
func NewCustomerController() *controllers.CustomerController {
	return &controllers.CustomerController{
		CreateCustomerUC:    NewCreateCustomerUC(),
		CreateWorkOrderUC:   NewCreateWorkOrderUC(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		UpdateCustomerUC:    NewUpdateCustomerUC(),
		Logger:              NewLogrusLogger(),
	}
}

// NewWorkOrderController función para inicializar el controlador de las ordenes de servicio
func NewWorkOrderController() *controllers.WorkOrderController {
	return &controllers.WorkOrderController{
		FinishWorkOrderUC:   NewFinishWorkOrderUC(),
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		UpdateWorkOrderUC:   NewUpdateWorkOrderUC(),
		Logger:              NewLogrusLogger(),
	}
}
