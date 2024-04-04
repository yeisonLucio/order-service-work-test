package factories

import "lucio.com/order-service/src/infra/controllers"

func NewCustomerController() *controllers.CustomerController {
	return &controllers.CustomerController{
		CreateCustomerUC:    NewCreateCustomerUC(),
		CreateWorkOrderUC:   NewCreateWorkOrderUC(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		UpdateCustomerUC:    NewUpdateCustomerUC(),
	}
}

func NewWorkOrderController() *controllers.WorkOrderController {
	return &controllers.WorkOrderController{
		FinishWorkOrderUC:   NewFinishWorkOrderUC(),
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		UpdateWorkOrderUC:   NewUpdateWorkOrderUC(),
	}
}
