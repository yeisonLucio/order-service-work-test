package factories

import (
	"lucio.com/order-service/src/application/customer"
	workorder "lucio.com/order-service/src/application/work-order"
	ICustomerUC "lucio.com/order-service/src/domain/customer/usecases"
	IWorkOrderUC "lucio.com/order-service/src/domain/workorder/usecases"
)

func NewCreateCustomerUC() ICustomerUC.CreateCustomerUC {
	return &customer.CreateCustomerUC{
		CustomerRepository: NewPostgresCustomerRepository(),
	}
}

func NewCreateWorkOrderUC() IWorkOrderUC.CreateWorkOrderUC {
	return &workorder.CreateWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		Time:                NewTimeLib(),
	}
}

func NewFinishWorkOrderUC() IWorkOrderUC.FinishWorkOrderUC {
	return &workorder.FinishWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		EventRepository:     NewRedisEventRepository(),
		Time:                NewTimeLib(),
	}
}

func NewUpdateCustomerUC() ICustomerUC.UpdateCustomerUC {
	return &customer.UpdateCustomerUC{
		CustomerRepository: NewPostgresCustomerRepository(),
	}
}

func NewUpdateWorkOrderUC() IWorkOrderUC.UpdateWorkOrderUC {
	return &workorder.UpdateWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
	}
}
