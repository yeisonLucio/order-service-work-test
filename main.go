package main

import (
	"github.com/joho/godotenv"
	"lucio.com/order-service/src"
	"lucio.com/order-service/src/config/postgres"
	"lucio.com/order-service/src/config/redis"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	postgres.Connect()
	postgres.RunMigrations()
	redis.Connect()

	app := src.GetApp()
	app.Run(":8080")
}
