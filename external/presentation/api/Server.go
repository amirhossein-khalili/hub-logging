package api

import (
	"log"

	"hub_logging/config"
	"hub_logging/external/infra/di"
	"hub_logging/external/presentation/api/rest"
	"hub_logging/external/presentation/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

// StartServer initializes Fiber, loads dependencies, registers routes, and starts the HTTP server.
func StartServer(cfg config.AppConfig) {
	// Initialize a new Fiber instance.
	app := fiber.New()

	// Initialize the dependency injection container.
	// This creates the DB connection, repositories, and use cases.
	container := di.NewContainer()

	// Create a REST handler instance (a wrapper to hold the Fiber app).
	restHandler := &rest.RestHandler{App: app}

	// Register LogMessage CRUD routes.
	handlers.SetupLogRoutes(restHandler, container.CreateLogUseCase, container.LogMessageRepo)

	// Start the server.
	log.Fatal(app.Listen(cfg.ServerPort))
}
