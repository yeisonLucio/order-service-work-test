package src

import (
	"github.com/gin-gonic/gin"
)

func GetApp() *gin.Engine {
	app := gin.Default()

	app = getRoutes(app)

	return app
}
