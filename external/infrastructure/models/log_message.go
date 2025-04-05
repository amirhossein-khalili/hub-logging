package models

import (
	"time"

	"github.com/google/uuid"
)

// LogMessage defines the structure of the log message table.
type LogMessage struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Timestamp    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"timestamp"`
	StatusCode   int       `gorm:"not null" json:"status_code"`
	HttpMethod   string    `gorm:"not null" json:"http_method"`
	RoutePath    string    `gorm:"not null" json:"route_path"`
	Message      string    `gorm:"type:text;not null" json:"message"`
	UserName     string    `gorm:"not null" json:"user_name"`
	DestHostname string    `gorm:"not null" json:"dest_hostname"`
	SourceIP     string    `gorm:"not null" json:"source_ip"`
}
