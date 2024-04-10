package repositories

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/dtos"
	"lucio.com/order-service/src/domain/workorder/entities"
)

const workOrderUpdatedStream string = "work-order-updated"

// RedisEventRepository estructura para enviar que permite acceder a métodos para enviar eventos
type RedisEventRepository struct {
	RedisClient *redis.Client
	Logger      *logrus.Logger
}

// NotifyWorkOrderFinished permite enviar mensaje por el stream work-order-updated
func (r *RedisEventRepository) NotifyWorkOrderFinished(payload *entities.WorkOrder) *dtos.CustomError {
	log := r.Logger.WithFields(logrus.Fields{
		"file":    "postgres_customer_repository",
		"method":  "FindByID",
		"payload": payload,
	})

	object, err := json.Marshal(payload)
	if err != nil {
		log.WithError(err).Error()
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
		log.WithError(err).Error()
		return &dtos.CustomError{
			Code:  500,
			Error: result.Err(),
		}
	}

	return nil
}
