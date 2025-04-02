package dtos

type CreateLogDTO struct {
	StatusCode   int
	HttpMethod   string
	RoutePath    string
	Message      string
	UserName     string
	DestHostname string
	SourceIP     string
}
