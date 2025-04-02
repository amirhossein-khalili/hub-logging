package entities

import "github.com/google/uuid"

// Statistics represents aggregated metrics for log entries.
type Statistics struct {
	ID           uuid.UUID `json:"id"`
	RoutePath    string    `json:"route_path"`
	StatusCode   int       `json:"status_code"`
	HttpMethod   string    `json:"http_method"`
	SourceIP     string    `json:"source_ip"`
	UserName     string    `json:"user_name"`
	RequestCount int       `json:"request_count"`
}
