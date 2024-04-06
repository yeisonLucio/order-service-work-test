package factories

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/helpers"
)

func NewTimeLib() helpers.Timer {
	return &helpers.DefaultTimer{}
}

func NewLogrusLogger() *logrus.Logger {
	return &logrus.Logger{}
}
