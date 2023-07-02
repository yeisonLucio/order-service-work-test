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
			customers.PUT("/:id", di.Container.CustomerController.UpdateCustomer)
			customers.GET("/:id", di.Container.CustomerController.GetCustomer)
			customers.DELETE("/:id", di.Container.CustomerController.DeleteCustomer)
			customers.POST("/:id/work-orders", di.Container.CustomerController.CreateWorkOrder)
			customers.GET("/:id/work-orders", di.Container.CustomerController.GetWorkOrders)
		}

		orders := api.Group("work-orders")
		{
			orders.GET("/", di.Container.WorkOrderController.GetWorkOrders)
			orders.PUT("/:id", di.Container.WorkOrderController.UpdateWorkOrder)
			orders.GET("/:id", di.Container.WorkOrderController.GetWorkOrder)
			orders.DELETE("/:id", di.Container.WorkOrderController.DeleteWorkOrder)
			orders.PATCH(":id/finish", di.Container.WorkOrderController.FinishWorkOrder)
		}
	}

	return app
}
