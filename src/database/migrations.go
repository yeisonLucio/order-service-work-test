package database

import "lucio.com/order-service/src/models"

func RunMigrations() (err error) {
	err = DB.AutoMigrate(&models.Customer{}, &models.WorkOrder{})
	return
}
