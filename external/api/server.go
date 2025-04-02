package api

import (
	"hub_logging/configs"
	"hub_logging/external/api/rest"
	"hub_logging/external/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config configs.AppConfig) {

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
