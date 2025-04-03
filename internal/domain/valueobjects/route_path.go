package valueobjects

import (
	"fmt"
	"strings"
)

// RoutePath represents a route path in the application.
type RoutePath string

// NewRoutePath creates a new RoutePath and validates it.
func NewRoutePath(path string) (RoutePath, error) {
	if !strings.HasPrefix(path, "/") {
		return "", fmt.Errorf("route path must start with '/': %s", path)
	}
	return RoutePath(path), nil
}

// String returns the string representation of the RoutePath.
func (rp RoutePath) String() string {
	return string(rp)
}
