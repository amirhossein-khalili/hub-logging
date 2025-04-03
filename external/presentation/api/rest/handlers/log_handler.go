package handlers

import (
	"fmt"
	"net/http"
	"time"

	"hub_logging/external/presentation/api/rest"
	"hub_logging/internal/application/dtos"
	"hub_logging/internal/application/usecases"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/gofiber/fiber/v2"
)

// LogHandler aggregates the use case and repository dependencies.
type LogHandler struct {
	// UseCase to create a log.
	CreateLogUseCase *usecases.CreateLogUseCase
	// Repository for direct LogMessage operations.
	LogMessageRepo repositoriesInterfaces.ILogMessageRepository
}

// SetupLogRoutes registers CRUD endpoints for logs.
func SetupLogRoutes(rh *rest.RestHandler, createLogUseCase *usecases.CreateLogUseCase, logRepo repositoriesInterfaces.ILogMessageRepository) {
	handler := LogHandler{
		CreateLogUseCase: createLogUseCase,
		LogMessageRepo:   logRepo,
	}
	app := rh.App

	// CRUD routes for logs.
	app.Get("/logs", handler.ListLogs)
	app.Get("/logs/:id", handler.GetLog)
	app.Post("/logs", handler.CreateLog)
	app.Put("/logs/:id", handler.UpdateLog)
	app.Delete("/logs/:id", handler.DeleteLog)
}

// ListLogs handles GET /logs and returns all log messages.
func (h *LogHandler) ListLogs(ctx *fiber.Ctx) error {
	logs, err := h.LogMessageRepo.FindAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(logs)
}

// GetLog handles GET /logs/:id and returns a specific log message.
func (h *LogHandler) GetLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	logMessage, err := h.LogMessageRepo.FindByID(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Log not found"})
	}
	return ctx.JSON(logMessage)
}

// CreateLog handles POST /logs to create a new log message.
func (h *LogHandler) CreateLog(ctx *fiber.Ctx) error {
	var req dtos.CreateLogDTO
	if err := ctx.BodyParser(&req); err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// If Timestamp is not provided, use the current time.
	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}

	// Use the use case to create a new log.
	if err := h.CreateLogUseCase.Execute(req); err != nil {
		fmt.Println(err)

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Log created successfully"})
}

// UpdateLogRequest represents the expected payload for updating a log.
type UpdateLogRequest struct {
	StatusCode   int    `json:"status_code"`
	HttpMethod   string `json:"http_method"`
	RoutePath    string `json:"route_path"`
	Message      string `json:"message"`
	UserName     string `json:"user_name"`
	DestHostname string `json:"dest_hostname"`
	SourceIP     string `json:"source_ip"`
}

// UpdateLog handles PUT /logs/:id to update an existing log message.
func (h *LogHandler) UpdateLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	// Retrieve the existing log.
	existing, err := h.LogMessageRepo.FindByID(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Log not found"})
	}

	var req UpdateLogRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// Update the fields.
	existing.StatusCode = req.StatusCode
	existing.HttpMethod = req.HttpMethod
	existing.RoutePath = req.RoutePath
	existing.Message = req.Message
	existing.UserName = req.UserName
	existing.DestHostname = req.DestHostname
	existing.SourceIP = req.SourceIP

	if err := h.LogMessageRepo.Update(existing); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Log updated successfully"})
}

// DeleteLog handles DELETE /logs/:id to remove a log message.
func (h *LogHandler) DeleteLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := h.LogMessageRepo.Delete(id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Log deleted successfully"})
}
