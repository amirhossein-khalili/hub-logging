package controller

import (
	"net/http"
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

func (lc *LogController) ListLogs(ctx *fiber.Ctx) error {
	logs, err := lc.LogMessageRepo.FindAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
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

type UpdateLogRequest struct {
	StatusCode   int    `json:"status_code"`
	HttpMethod   string `json:"http_method"`
	RoutePath    string `json:"route_path"`
	Message      string `json:"message"`
	UserName     string `json:"user_name"`
	DestHostname string `json:"dest_hostname"`
	SourceIP     string `json:"source_ip"`
}

func (lc *LogController) UpdateLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	// Retrieve the existing log.
	existing, err := lc.LogMessageRepo.FindByID(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Log not found"})
	}

	var req UpdateLogRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// Update fields.
	existing.StatusCode = req.StatusCode
	existing.HttpMethod = req.HttpMethod
	existing.RoutePath = req.RoutePath
	existing.Message = req.Message
	existing.UserName = req.UserName
	existing.DestHostname = req.DestHostname
	existing.SourceIP = req.SourceIP

	if err := lc.LogMessageRepo.Update(existing); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Log updated successfully"})
}

func (lc *LogController) DeleteLog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := lc.LogMessageRepo.Delete(id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Log deleted successfully"})
}
