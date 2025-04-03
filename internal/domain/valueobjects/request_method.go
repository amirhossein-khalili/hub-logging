package valueobjects

import (
	"fmt"
	"strings"
)

// RequestMethod represents an HTTP request method.
type RequestMethod string

// Valid HTTP methods
const (
	GET     RequestMethod = "GET"
	POST    RequestMethod = "POST"
	PUT     RequestMethod = "PUT"
	DELETE  RequestMethod = "DELETE"
	PATCH   RequestMethod = "PATCH"
	OPTIONS RequestMethod = "OPTIONS"
	HEAD    RequestMethod = "HEAD"
)

// NewRequestMethod creates a new RequestMethod and validates it.
func NewRequestMethod(method string) (RequestMethod, error) {
	switch strings.ToUpper(method) {
	case "GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD":
		return RequestMethod(strings.ToUpper(method)), nil
	default:
		return "", fmt.Errorf("invalid request method: %s", method)
	}
}

// String returns the string representation of the RequestMethod.
func (rm RequestMethod) String() string {
	return string(rm)
}
