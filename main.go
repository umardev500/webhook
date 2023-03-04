package main

import (
	"fmt"
	"log"
	"os"
	"webhook/config"
	"webhook/injector"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var conn config.Connection
	order := conn.OrderConn()
	customer := conn.CustomerConn()

	port := os.Getenv("PORT")
	app := fiber.New()
	injector.NewOrderInjector(app, order, customer)

	fmt.Printf("⚡️[server]: Server is running on port %s\n", port)
	log.Fatal(app.Listen(port))
}
