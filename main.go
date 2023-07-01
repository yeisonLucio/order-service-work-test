package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"lucio.com/order-service/src"
	"lucio.com/order-service/src/database"
	"lucio.com/order-service/src/di"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if err := database.RunMigrations(); err != nil {
		fmt.Println("error corriendo las migraciones")
	}

	di.BuildContainer()

	app := src.GetApp()

	app.Run(":8080")
}
