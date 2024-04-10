package factories

import (
	"lucio.com/order-service/src/application/customer"
	workorder "lucio.com/order-service/src/application/work-order"
	ICustomerUC "lucio.com/order-service/src/domain/customer/usecases"
	IWorkOrderUC "lucio.com/order-service/src/domain/workorder/usecases"
)

// NewCreateCustomerUC función para inicializar el caso de uso de crear nuevo cliente
func NewCreateCustomerUC() ICustomerUC.CreateCustomerUC {
	return &customer.CreateCustomerUC{
		CustomerRepository: NewPostgresCustomerRepository(),
		Logger:             NewLogrusLogger(),
	}
}

// NewCreateWorkOrderUC función para inicializar el caso de uso de crear nueva orden de servicio
func NewCreateWorkOrderUC() IWorkOrderUC.CreateWorkOrderUC {
	return &workorder.CreateWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		Time:                NewTimeLib(),
		Logger:              NewLogrusLogger(),
	}
}

// NewFinishWorkOrderUC función para inicializar el caso de uso de finalizar una orden de servicio
func NewFinishWorkOrderUC() IWorkOrderUC.FinishWorkOrderUC {
	return &workorder.FinishWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		CustomerRepository:  NewPostgresCustomerRepository(),
		EventRepository:     NewRedisEventRepository(),
		Time:                NewTimeLib(),
		Logger:              NewLogrusLogger(),
	}
}

// NewUpdateCustomerUC función para inicializar el caso de uso de actualizar cliente
func NewUpdateCustomerUC() ICustomerUC.UpdateCustomerUC {
	return &customer.UpdateCustomerUC{
		CustomerRepository: NewPostgresCustomerRepository(),
		Logger:             NewLogrusLogger(),
	}
}

// NewUpdateWorkOrderUC función para inicializar el caso de uso de actualizar orden de servicio
func NewUpdateWorkOrderUC() IWorkOrderUC.UpdateWorkOrderUC {
	return &workorder.UpdateWorkOrderUC{
		WorkOrderRepository: NewPostgresWorkOrderRepository(),
		Logger:              NewLogrusLogger(),
	}
}
