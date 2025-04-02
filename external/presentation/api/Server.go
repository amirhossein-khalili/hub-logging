package api

import (
	"hub_logging/config"
	"hub_logging/external/presentation/api/rest"
	"hub_logging/external/presentation/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {

	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}
	SetupRoutes(rh)

	app.Listen(config.ServerPort)
}

func SetupRoutes(rh *rest.RestHandler) {
	// log handlers
	handlers.SetupLogRoute(rh)

}
