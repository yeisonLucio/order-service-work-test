package src

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lucio.com/order-service/src/controllers"
	"lucio.com/order-service/src/database"
	"lucio.com/order-service/src/repositories"
	"lucio.com/order-service/src/usecases"
)

func GetApp() *gin.Engine {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := database.RunMigrations(); err != nil {
		fmt.Println("error corriendo las migraciones")
	}

	app := gin.Default()

	customerController := &controllers.CustomerController{
		CreateCustomerUC: &usecases.CreateCustomerUC{
			CustomerRepository: &repositories.PostgresCustomerRepository{
				ClientDB: database.DB,
			},
		},
	}

	app.POST("api/v1/customers", customerController.CreateCustomer)

	return app
}
