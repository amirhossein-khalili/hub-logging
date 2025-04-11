package rest

import (
	"hub_logging/external/presentation/api/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRestRoutes(app *fiber.App, statsController *controller.StatisticsController, logController *controller.LogController) {
	api := app.Group("/api/v1")
	stats := api.Group("/stats")

	/*----------------------------------------------------
	 *
	 *		STATS END POINTS
	 *
	 ----------------------------------------------------*/
	// Endpoint to get statistics (e.g., /api/v1/stats/ip?start=2025-04-10T00:00:00Z&end=2025-04-11T00:00:00Z)

	stats.Get("/ip", statsController.GetIPStatistics)
	stats.Get("/route", statsController.GetRouteStatistics)
	stats.Get("/method", statsController.GetMethodStatistics)
	stats.Get("/user", statsController.GetUserStatistics)

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
