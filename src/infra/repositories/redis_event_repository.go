package repositories

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

const workOrderUpdatedStream string = "work-order-updated"

type RedisEventRepository struct {
	RedisClient *redis.Client
}

func (r *RedisEventRepository) NotifyWorkOrderFinished(payload *entities.WorkOrder) *dtos.CustomError {
	object, err := json.Marshal(payload)
	if err != nil {
		return &dtos.CustomError{
			Code:  500,
			Error: err,
		}
	}

	result := r.RedisClient.XAdd(
		&redis.XAddArgs{
			ID:     "*",
			Stream: workOrderUpdatedStream,
			Values: map[string]interface{}{
				"data": object,
			},
		},
	)

	if result.Err() != nil {
		return &dtos.CustomError{
			Code:  500,
			Error: result.Err(),
		}
	}

	return nil
}
