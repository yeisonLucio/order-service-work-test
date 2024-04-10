package postgres

import (
	"fmt"

	"lucio.com/order-service/src/infra/models"
)

// RunMigrations migra los modelos de base de datos
func RunMigrations() {
	err := DB.AutoMigrate(
		&models.Customer{},
		&models.WorkOrder{},
	)

	if err != nil {
		fmt.Println("error corriendo las migraciones")
	}
}
