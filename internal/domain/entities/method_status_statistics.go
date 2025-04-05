package entities

import (
	"hub_logging/internal/domain/valueobjects"
	"time"

	"github.com/google/uuid"
)

type MethodStatusStatistics struct {
	ID            uuid.UUID                  `json:"id"`
	HttpMethod    valueobjects.RequestMethod `json:"http_method"`
	PeriodStart   time.Time                  `json:"period_start"`
	PeriodEnd     time.Time                  `json:"period_end"`
	TotalRequests int                        `json:"total_requests"`
	SuccessCount  int                        `json:"success_count"`
	ErrorCount    int                        `json:"error_count"`
}
