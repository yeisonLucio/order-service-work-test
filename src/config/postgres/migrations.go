package postgres

import (
	"fmt"

	"lucio.com/order-service/src/models"
)

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.Customer{},
		&models.WorkOrder{},
	)

	if err != nil {
		fmt.Println("error corriendo las migraciones")
	}
}
