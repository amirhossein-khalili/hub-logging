package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserStatistics struct {
	ID            uuid.UUID `json:"id"`
	UserName      string    `json:"user_name"`
	PeriodStart   time.Time `json:"period_start"`
	PeriodEnd     time.Time `json:"period_end"`
	TotalRequests int       `json:"total_requests"`
}
