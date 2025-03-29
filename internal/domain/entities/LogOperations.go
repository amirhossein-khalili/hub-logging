package entities

import "time"

// LogOperations represents metadata about log processing operations.
type LogOperations struct {
	ID           int       `json:"id"`
	Operation    string    `json:"operation"` // e.g., "create", "update", "delete"
	Timestamp    time.Time `json:"timestamp"`
	LogMessageID int       `json:"log_message_id"` // References a LogMessage
	PerformedBy  string    `json:"performed_by"`   // User or system that performed the operation
}
