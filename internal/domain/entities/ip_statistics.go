package entities

import (
	"time"

	"github.com/google/uuid"
)

type IPStatistics struct {
	ID            uuid.UUID `json:"id"`
	SourceIP      string    `json:"source_ip"`
	PeriodStart   time.Time `json:"period_start"`
	PeriodEnd     time.Time `json:"period_end"`
	TotalRequests int       `json:"total_requests"`
}
