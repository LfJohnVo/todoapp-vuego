package main

import (
	"backend/config"
	"backend/databases"
	"backend/middlewares"
	"backend/pkg/utils"
	"backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config.
	config := config.FiberConfig()

	//start fiber
	app := fiber.New(config)

	// Middlewares.
	middlewares.Params(app)

	//connect database
	databases.Connect()

	routes.SetupRoutes(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}
