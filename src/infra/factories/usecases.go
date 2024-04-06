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
		Logger:             NewLogrusLogger(),
	}
}

func NewCreateWorkOrderUC() IWorkOrderUC.CreateWorkOrderUC {
	return &workorder.CreateWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		Time:                NewTimeLib(),
		Logger:              NewLogrusLogger(),
	}
}

func NewFinishWorkOrderUC() IWorkOrderUC.FinishWorkOrderUC {
	return &workorder.FinishWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		EventRepository:     NewRedisEventRepository(),
		Time:                NewTimeLib(),
		Logger:              NewLogrusLogger(),
	}
}

func NewUpdateCustomerUC() ICustomerUC.UpdateCustomerUC {
	return &customer.UpdateCustomerUC{
		CustomerRepository: NewPostgresCustomerRepository(),
		Logger:             NewLogrusLogger(),
	}
}

func NewUpdateWorkOrderUC() IWorkOrderUC.UpdateWorkOrderUC {
	return &workorder.UpdateWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		Logger:              NewLogrusLogger(),
	}
}
