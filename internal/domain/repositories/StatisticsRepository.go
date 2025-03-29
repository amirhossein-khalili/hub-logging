package repositories

import "hub_logging/internal/domain/entities"

// StatisticsRepository defines the interface for Statistics persistence.
type StatisticsRepository interface {
	// Save persists a single Statistics entity.
	Save(statistics entities.Statistics) error

	// FindByRoutePath retrieves Statistics for a specific route path.
	FindByRoutePath(routePath string) ([]entities.Statistics, error)

	// FindByStatusCode retrieves Statistics for a specific HTTP status code.
	FindByStatusCode(statusCode int) ([]entities.Statistics, error)
}
