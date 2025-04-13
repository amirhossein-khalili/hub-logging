package repositoriesInterfaces

import (
	"context"
	"hub_logging/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type IRouteStatisticsRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.RouteStatistics, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.RouteStatistics, error)
	Create(ctx context.Context, stats *entities.RouteStatistics) error
	Update(ctx context.Context, stats *entities.RouteStatistics) error
}
