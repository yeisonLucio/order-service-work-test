package repositories

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"lucio.com/order-service/src/models"
)

const workOrderUpdatedStream string = "work-order-updated"

type RedisEventRepository struct {
	RedisClient *redis.Client
}

func (r *RedisEventRepository) NotifyWorkOrderFinished(workOrder models.WorkOrder) error {
	object, err := json.Marshal(workOrder)
	if err != nil {
		return err
	}

	result := r.RedisClient.XAdd(
		context.Background(),
		&redis.XAddArgs{
			ID:     "*",
			Stream: workOrderUpdatedStream,
			Values: map[string]interface{}{
				"data": object,
			},
		},
	)

	if result.Err() != nil {
		return err
	}

	return nil
}
