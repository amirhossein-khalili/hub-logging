package repositoriesInterfaces

import (
	"context"
	"github.com/google/uuid"
	"hub_logging/internal/domain/entities"
	"time"
)

type IMethodStatusStatisticsRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.MethodStatusStatistics, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.MethodStatusStatistics, error)
	Create(ctx context.Context, stats *entities.MethodStatusStatistics) error
	Update(ctx context.Context, stats *entities.MethodStatusStatistics) error
}
