package valueobjects

import "fmt"

// StatusCode represents an HTTP status code.
type StatusCode int

// NewStatusCode creates a new StatusCode and validates it.
func NewStatusCode(code int) (StatusCode, error) {
	if code < 100 || code > 599 {
		return 0, fmt.Errorf("invalid status code: %d", code)
	}
	return StatusCode(code), nil
}

// Int returns the integer value of the StatusCode.
func (sc StatusCode) Int() int {
	return int(sc)
}
