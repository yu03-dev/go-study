package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yu03-dev/go-study-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFuct)

	app.Post("/fact", handlers.CreateFact)
}
