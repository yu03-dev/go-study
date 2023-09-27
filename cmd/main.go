package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yu03-dev/go-study-api/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
