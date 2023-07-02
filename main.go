package main

import (
	"github.com/joho/godotenv"
	"lucio.com/order-service/src"
	"lucio.com/order-service/src/config/postgres"
	"lucio.com/order-service/src/config/redis"
	"lucio.com/order-service/src/di"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	postgres.Connect()
	postgres.RunMigrations()
	redis.Connect()
	di.BuildContainer()

	app := src.GetApp()

	app.Run(":8080")
}
