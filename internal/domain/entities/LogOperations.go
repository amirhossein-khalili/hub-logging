package entities

import (
	"time"

	"github.com/google/uuid"
)

// LogOperations represents metadata about log processing operations.
type LogOperations struct {
	ID           uuid.UUID `json:"id"`
	Operation    string    `json:"operation"` // e.g., "create", "update", "delete"
	Timestamp    time.Time `json:"timestamp"`
	LogMessageID uuid.UUID `json:"log_message_id"` // References a LogMessage
	PerformedBy  string    `json:"performed_by"`   // User or system that performed the operation
}
