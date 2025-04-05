package rest

import (
	"hub_logging/external/presentation/api/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRestRoutes(app *fiber.App, logController *controller.LogController) {
	api := app.Group("/api/v1")

	/*----------------------------------------------------
	 *
	 *		CRUD OPERATIONS FOR LOGS
	 *
	 ----------------------------------------------------*/
	logs := api.Group("/logs")
	logs.Get("/", logController.ListLogs)
	logs.Get("/:id", logController.GetLog)
	logs.Post("/", logController.CreateLog)
	logs.Delete("/:id", logController.DeleteLog)
}
