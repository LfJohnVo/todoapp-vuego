package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {

	home := app.Group("/home")
	home.Static("/", "./public/index.html")

	// Middleware
	api := app.Group("/api")

	//Index endpoint
	api.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	//todo
	api.Get("/tasks", handlers.GetTasks)
	api.Get("/tasks/:id", handlers.GetTask)
	api.Post("/tasks", handlers.AddTask)
	api.Put("/tasks", handlers.UpdateTask)
	api.Delete("tasks/:id", handlers.RemoveTask)
}
