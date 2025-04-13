package controller

import (
	"net/http"
	"strconv"
	"time"

	"hub_logging/internal/application/dtos"
	"hub_logging/internal/application/usecases"

	"github.com/gofiber/fiber/v2"
)

type LogController struct {
	CreateLogUseCase *usecases.CreateLogUseCase
	DeleteLogUseCase *usecases.DeleteLogUseCase
	GetLogsUseCase   *usecases.GetLogsUseCase
}

func NewLogController(createLogUseCase *usecases.CreateLogUseCase, deleteLogUseCase *usecases.DeleteLogUseCase, getLogsUseCase *usecases.GetLogsUseCase) *LogController {
	return &LogController{
		CreateLogUseCase: createLogUseCase,
		DeleteLogUseCase: deleteLogUseCase,
		GetLogsUseCase:   getLogsUseCase,
	}
}

// ListLogs handles GET /logs and returns a paginated list of log messages.
func (lc *LogController) ListLogs(ctx *fiber.Ctx) error {
	page := 1
	limit := 10

	// Extracting context from the request
	requestCtx := ctx.Context()

	// Parse query parameters for pagination
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

	offset := (page - 1) * limit

	logs, err := lc.GetLogsUseCase.Execute(requestCtx, limit, offset) // Pass the context here
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(logs)
}

func (lc *LogController) GetLog(ctx *fiber.Ctx) error {
	// Extracting context from the request
	requestCtx := ctx.Context()

	id := ctx.Params("id")
	logMessage, err := lc.GetLogsUseCase.ExecuteSingle(requestCtx, id) // Pass the context here
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Log not found"})
	}
	return ctx.JSON(logMessage)
}

func (lc *LogController) CreateLog(ctx *fiber.Ctx) error {
	requestCtx := ctx.Context()

	var req dtos.CreateLogDTO
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// If Timestamp is not provided, use the current time.
	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now()
	}

	if err := lc.CreateLogUseCase.Execute(requestCtx, req); err != nil { // Pass the context here
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "Log created successfully"})
}

func (lc *LogController) DeleteLog(ctx *fiber.Ctx) error {
	requestCtx := ctx.Context()

	id := ctx.Params("id")
	if err := lc.DeleteLogUseCase.Execute(requestCtx, id); err != nil { // Pass the context here
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Log deleted successfully"})
}
