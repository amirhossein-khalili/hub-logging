package server

import (
	"log"

	"hub_logging/config"
	"hub_logging/external/infra/di"
	"hub_logging/external/presentation/api/controller"
	"hub_logging/external/presentation/api/middleware"
	"hub_logging/external/presentation/api/routes/rest"

	"github.com/gofiber/fiber/v2"
)

// StartServer initializes the Fiber app, sets up dependencies, registers routes,
// and starts the HTTP server.
func StartServer(cfg config.AppConfig) {
	// Initialize a new Fiber instance.
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorMiddleware,
	})

	// Initialize the container with dependencies.
	container, err := di.InitializeContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
	// Create an instance of LogController using the dependencies provided by the container.
	logController := controller.NewLogController(container.CreateLogUseCase, container.LogMessageRepo)

	// Register the REST routes (grouping under /api, for example).
	rest.RegisterRestRoutes(app, logController)

	// Start the Fiber server on the port specified in configuration.
	log.Fatal(app.Listen(cfg.ServerPort))
}
