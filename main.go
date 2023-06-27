package main

import "lucio.com/order-service/src"

func main() {
	app := src.GetApp()

	app.Run(":8080")
}
