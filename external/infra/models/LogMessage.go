package models

import (
	"time"

	"github.com/google/uuid"
)

type LogMessage struct {
	ID           uuid.UUID `gorm:"primary_key"`
	Timestamp    time.Time
	StatusCode   int
	HttpMethod   string
	RoutePath    string
	Message      string
	UserName     string
	DestHostname string
	SourceIP     string
}
