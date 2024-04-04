package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "lucio.com/order-service/docs"
	"lucio.com/order-service/src/infra/factories"
)

func GetRoutes(app *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT"))
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	customerController := factories.NewCustomerController()
	workOrderController := factories.NewWorkOrderController()

	api := app.Group("api/v1")
	{
		customers := api.Group("customers")
		{
			customers.POST("/", customerController.CreateCustomer)
			customers.GET("/", customerController.GetCustomers)
			customers.PUT("/:id", customerController.UpdateCustomer)
			customers.GET("/:id", customerController.GetCustomer)
			customers.DELETE("/:id", customerController.DeleteCustomer)
			customers.POST("/:id/work-orders", customerController.CreateWorkOrder)
			customers.GET("/:id/work-orders", customerController.GetWorkOrders)
		}

		orders := api.Group("work-orders")
		{
			orders.GET("/", workOrderController.GetWorkOrders)
			orders.PUT("/:id", workOrderController.UpdateWorkOrder)
			orders.GET("/:id", workOrderController.GetWorkOrder)
			orders.DELETE("/:id", workOrderController.DeleteWorkOrder)
			orders.PATCH(":id/finish", workOrderController.FinishWorkOrder)
		}
	}

	return app
}
