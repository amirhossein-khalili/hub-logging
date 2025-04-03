package dtos

import "time"

type CreateLogDTO struct {
	Timestamp    time.Time `json:"timestamp"`
	StatusCode   int       `json:"status_code"`
	HttpMethod   string    `json:"http_method"`
	RoutePath    string    `json:"route_path"`
	Message      string    `json:"message"`
	UserName     string    `json:"user_name"`
	DestHostname string    `json:"dest_hostname"`
	SourceIP     string    `json:"source_ip"`
}
