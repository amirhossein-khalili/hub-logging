package repositoriesInterfaces

import (
	"context"
	"hub_logging/internal/domain/entities"
)

// IStatisticsRepository defines the interface for Statistics persistence.
type IStatisticsRepository interface {
	Save(ctx context.Context, statistics entities.Statistics) error
	FindByRoutePath(ctx context.Context, routePath string) ([]entities.Statistics, error)
	FindByStatusCode(ctx context.Context, statusCode int) ([]entities.Statistics, error)
}
