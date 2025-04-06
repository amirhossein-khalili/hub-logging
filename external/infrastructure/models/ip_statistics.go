package models

import (
	"time"

	"github.com/google/uuid"
)

type IPStatistics struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	SourceIP      string    `gorm:"type:varchar(45);not null"`
	PeriodStart   time.Time `gorm:"type:timestamp;not null"`
	PeriodEnd     time.Time `gorm:"type:timestamp;not null"`
	TotalRequests int       `gorm:"type:integer;not null"`
}
