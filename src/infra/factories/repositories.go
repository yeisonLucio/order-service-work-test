package factories

import (
	"lucio.com/order-service/src/config/postgres"
	"lucio.com/order-service/src/config/redis"
	ICustomerRepos "lucio.com/order-service/src/domain/customer/repositories"
	IWorkOrderRepos "lucio.com/order-service/src/domain/workorder/repositories"

	"lucio.com/order-service/src/infra/repositories"
)

func NewPostgresCustomerRepository() ICustomerRepos.CustomerRepository {
	return &repositories.PostgresCustomerRepository{
		ClientDB: postgres.DB,
		Logger:   NewLogrusLogger(),
	}
}

func NewPostgresWorkOrderRepository() IWorkOrderRepos.WorkOrderRepository {
	return &repositories.PostgresWorkOrderRepository{
		ClientDB: postgres.DB,
		Logger:   NewLogrusLogger(),
	}
}

func NewRedisEventRepository() IWorkOrderRepos.EventRepository {
	return &repositories.RedisEventRepository{
		RedisClient: redis.RedisClient,
		Logger:      NewLogrusLogger(),
	}
}
