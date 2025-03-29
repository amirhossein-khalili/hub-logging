package entities

import "time"

// LogMessage represents a single log entry in the system.
type LogMessage struct {
	ID           int       `json:"id"`
	Timestamp    time.Time `json:"timestamp"`
	StatusCode   int       `json:"status_code"`
	HttpMethod   string    `json:"http_method"`
	RoutePath    string    `json:"route_path"`
	Message      string    `json:"message"`
	UserName     string    `json:"user_name"`
	DestHostname string    `json:"dest_hostname"`
	SourceIP     string    `json:"source_ip"`
}
