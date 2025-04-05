package controller

import (
	"net/http"
	"strconv"
	"time"

	"hub_logging/internal/application/dtos"
	"hub_logging/internal/application/usecases"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/gofiber/fiber/v2"
)

type LogController struct {
	CreateLogUseCase *usecases.CreateLogUseCase
	LogMessageRepo   repositoriesInterfaces.ILogMessageRepository
}

func NewLogController(createLogUseCase *usecases.CreateLogUseCase, repo repositoriesInterfaces.ILogMessageRepository) *LogController {
	return &LogController{
		CreateLogUseCase: createLogUseCase,
		LogMessageRepo:   repo,
	}
}

// ListLogs handles GET /logs and returns a paginated list of log messages.
func (lc *LogController) ListLogs(ctx *fiber.Ctx) error {
	// Get pagination parameters from query parameters, with default values.
	page := 1
	limit := 10

	// Parse page and limit from query parameters, if provided.
	if p := ctx.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}
	if l := ctx.Query("limit"); l != "" {
		if parsedLimit, err := strconv.Atoi(l); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Calculate the offset.
	offset := (page - 1) * limit

	// Fetch logs with pagination.
	logs, err := lc.LogMessageRepo.FindWithPagination(limit, offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the paginated logs.
	return ctx.JSON(logs)
}

func (lc *LogController) GetLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	logMessage, err := lc.LogMessageRepo.FindByID(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Log not found"})
	}
	return ctx.JSON(logMessage)
}

func (lc *LogController) CreateLog(ctx *fiber.Ctx) error {
	var req dtos.CreateLogDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// If Timestamp is not provided, use the current time.
	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}

	// Execute the use case to create a new log.
	if err := lc.CreateLogUseCase.Execute(req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Log created successfully"})
}

func (lc *LogController) DeleteLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := lc.LogMessageRepo.Delete(id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Log deleted successfully"})
}
