package controller

import (
	"context"
	"net/http"
	"time"

	"hub_logging/internal/application/usecases"

	"github.com/gofiber/fiber/v2"
)

// StatisticsController handles endpoints for aggregated statistics.
type StatisticsController struct {
	IPStatsUseCase             *usecases.GetIPStatisticsUseCase
	RouteStatsUseCase          *usecases.GetRouteStatisticsUseCase
	GetMethodStatisticsUseCase *usecases.GetMethodStatisticsUseCase
	GetUserStatisticsUseCase   *usecases.GetUserStatisticsUseCase
}

// NewStatisticsController creates a new StatisticsController.
func NewStatisticsController(ipUseCase *usecases.GetIPStatisticsUseCase, routeUseCase *usecases.GetRouteStatisticsUseCase, methodUseCase *usecases.GetMethodStatisticsUseCase, userUseCase *usecases.GetUserStatisticsUseCase) *StatisticsController {
	return &StatisticsController{
		IPStatsUseCase:             ipUseCase,
		RouteStatsUseCase:          routeUseCase,
		GetMethodStatisticsUseCase: methodUseCase,
		GetUserStatisticsUseCase:   userUseCase,
	}
}

// GetIPStatistics retrieves IP statistics.
// Optional query parameters: "start" and "end" in RFC3339 format.
func (sc *StatisticsController) GetIPStatistics(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start")
	endStr := ctx.Query("end")
	var start, end time.Time
	var err error

	if startStr == "" {
		start = time.Now().Truncate(24 * time.Hour)
	} else {
		start, err = time.Parse(time.RFC3339, startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
		}
	}

	if endStr == "" {
		end = start.Add(24 * time.Hour)
	} else {
		end, err = time.Parse(time.RFC3339, endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
		}
	}

	stats, err := sc.IPStatsUseCase.Execute(context.Background(), start, end)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stats)
}

// GetRouteStatistics retrieves route statistics.
// Optional query parameters: "start" and "end" in RFC3339 format.
func (sc *StatisticsController) GetRouteStatistics(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start")
	endStr := ctx.Query("end")
	var start, end time.Time
	var err error

	if startStr == "" {
		start = time.Now().Truncate(24 * time.Hour)
	} else {
		start, err = time.Parse(time.RFC3339, startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
		}
	}

	if endStr == "" {
		end = start.Add(24 * time.Hour)
	} else {
		end, err = time.Parse(time.RFC3339, endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
		}
	}

	stats, err := sc.RouteStatsUseCase.Execute(context.Background(), start, end)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stats)
}

// GetMethodStatistics retrieves Method Statistics
// Optional query parameters: "start" and "end" in RFC3339 format.
func (sc *StatisticsController) GetMethodStatistics(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start")
	endStr := ctx.Query("end")
	var start, end time.Time
	var err error

	if startStr == "" {
		start = time.Now().Truncate(24 * time.Hour)
	} else {
		start, err = time.Parse(time.RFC3339, startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
		}
	}

	if endStr == "" {
		end = start.Add(24 * time.Hour)
	} else {
		end, err = time.Parse(time.RFC3339, endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
		}
	}

	stats, err := sc.GetMethodStatisticsUseCase.Execute(context.Background(), start, end)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stats)
}

// GetUserStatistics retrieves User Statistics
// Optional query parameters: "start" and "end" in RFC3339 format.
func (sc *StatisticsController) GetUserStatistics(ctx *fiber.Ctx) error {
	startStr := ctx.Query("start")
	endStr := ctx.Query("end")
	var start, end time.Time
	var err error

	if startStr == "" {
		start = time.Now().Truncate(24 * time.Hour)
	} else {
		start, err = time.Parse(time.RFC3339, startStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date"})
		}
	}

	if endStr == "" {
		end = start.Add(24 * time.Hour)
	} else {
		end, err = time.Parse(time.RFC3339, endStr)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date"})
		}
	}

	stats, err := sc.GetUserStatisticsUseCase.Execute(context.Background(), start, end)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(stats)
}
