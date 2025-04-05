package entities

import (
    "time"
    "github.com/google/uuid"
	"hub_logging/internal/domain/valueobjects"
)

type RouteStatistics struct {
    ID            uuid.UUID          `json:"id"`
    RoutePath     valueobjects.RoutePath `json:"route_path"`
    PeriodStart   time.Time          `json:"period_start"`
    PeriodEnd     time.Time          `json:"period_end"`
    TotalRequests int                `json:"total_requests"`
    SuccessCount  int                `json:"success_count"`
    ErrorCount    int                `json:"error_count"`
}