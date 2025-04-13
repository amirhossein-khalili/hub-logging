package repositoriesInterfaces

import (
	"context"
	"hub_logging/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type IIPStatisticsRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.IPStatistics, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.IPStatistics, error)
	Create(ctx context.Context, stats *entities.IPStatistics) error
	Update(ctx context.Context, stats *entities.IPStatistics) error
}
