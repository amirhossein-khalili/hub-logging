package models

import (
	"hub_logging/internal/domain/valueobjects"
	"time"

	"github.com/google/uuid"
)

type RouteStatistics struct {
	ID            uuid.UUID              `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	RoutePath     valueobjects.RoutePath `gorm:"type:varchar(255);not null"`
	PeriodStart   time.Time              `gorm:"type:timestamp;not null"`
	PeriodEnd     time.Time              `gorm:"type:timestamp;not null"`
	TotalRequests int                    `gorm:"type:integer;not null"`
	SuccessCount  int                    `gorm:"type:integer;not null"`
	ErrorCount    int                    `gorm:"type:integer;not null"`
}
