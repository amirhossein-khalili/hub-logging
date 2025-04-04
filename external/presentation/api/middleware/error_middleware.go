package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ErrorMiddleware is a custom error handler middleware that logs errors
// and returns a JSON response with a consistent structure.
func ErrorMiddleware(c *fiber.Ctx, err error) error {
	// Log the error.
	log.Println("Error:", err)

	// Determine status code; if error is a Fiber error, use its code, else 500.
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Return JSON response.
	return c.Status(code).JSON(fiber.Map{
		"success":    false,
		"error":      err.Error(),
		"message":    err.Error(),
		"statusCode": code,
	})
}
