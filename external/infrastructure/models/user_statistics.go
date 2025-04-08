package models

import (
	"github.com/google/uuid"
	"time"
)

type UserStatistics struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserName      string    `gorm:"type:varchar(100);not null"`
	PeriodStart   time.Time `gorm:"type:timestamp;not null"`
	PeriodEnd     time.Time `gorm:"type:timestamp;not null"`
	TotalRequests int       `gorm:"type:integer;not null"`
}
