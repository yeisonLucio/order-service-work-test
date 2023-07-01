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
			customers.GET("/", di.Container.CustomerController.GetCustomers)
			customers.POST("/:id/work-orders", di.Container.CustomerController.CreateWorkOrder)
			customers.GET("/:id/work-orders", di.Container.CustomerController.GetWorkOrders)
		}

		orders := api.Group("work-orders")
		{
			orders.GET("/", di.Container.WorkOrderController.GetWorkOrders)
			orders.PATCH(":id/finish", di.Container.WorkOrderController.FinishWorkOrder)
		}
	}

	return app
}
