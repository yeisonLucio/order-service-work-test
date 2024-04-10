package src

import (
	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/infra/routes"
)

// GetApp genera una aplicación con sus rutas
func GetApp() *gin.Engine {
	app := gin.Default()

	app = routes.GetRoutes(app)

	return app
}
