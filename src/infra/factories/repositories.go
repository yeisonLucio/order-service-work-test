package factories

import (
	"lucio.com/order-service/src/config/postgres"
	"lucio.com/order-service/src/config/redis"
	ICustomerRepos "lucio.com/order-service/src/domain/customer/repositories"
	IWorkOrderRepos "lucio.com/order-service/src/domain/workorder/repositories"

	"lucio.com/order-service/src/infra/repositories"
)

// NewPostgresCustomerRepository función para inicializar repository de los clientes
func NewPostgresCustomerRepository() ICustomerRepos.CustomerRepository {
	return &repositories.PostgresCustomerRepository{
		ClientDB: postgres.DB,
		Logger:   NewLogrusLogger(),
	}
}

// NewPostgresWorkOrderRepository función para inicializar repository de las ordenes de servicio
func NewPostgresWorkOrderRepository() IWorkOrderRepos.WorkOrderRepository {
	return &repositories.PostgresWorkOrderRepository{
		ClientDB: postgres.DB,
		Logger:   NewLogrusLogger(),
	}
}

// NewRedisEventRepository función para inicializar repository de eventos
func NewRedisEventRepository() IWorkOrderRepos.EventRepository {
	return &repositories.RedisEventRepository{
		RedisClient: redis.RedisClient,
		Logger:      NewLogrusLogger(),
	}
}
