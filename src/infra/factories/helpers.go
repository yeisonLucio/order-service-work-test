package factories

import (
	"github.com/sirupsen/logrus"
	"lucio.com/order-service/src/domain/common/helpers"
)

// NewTimeLib función para inicializar helper de la librería time
func NewTimeLib() helpers.Timer {
	return &helpers.DefaultTimer{}
}

// NewLogrusLogger función para inicializar librería de logger
func NewLogrusLogger() *logrus.Logger {
	return &logrus.Logger{}
}
