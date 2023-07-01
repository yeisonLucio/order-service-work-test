package src

import (
	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/di"
)

func getRoutes(app *gin.Engine) *gin.Engine {

	api := app.Group("api/v1")
	{
		customers := api.Group("customers")
		{
			customers.POST("/", di.Container.CustomerController.CreateCustomer)
			customers.GET("/")
			customers.POST("/:id/work-orders")
			customers.GET("/:id/work-orders")
		}

		orders := api.Group("work-orders")
		{
			orders.GET("/")
			orders.PATCH(":id/finish")
		}
	}

	return app
}
