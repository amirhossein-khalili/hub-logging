package repositoriesInterfaces

import (
    "context"
    "time"
    "github.com/google/uuid"
    "hub_logging/internal/domain/entities"
)

type IMethodStatusStatisticsRepository interface {
    GetByID(ctx context.Context, id uuid.UUID) (*entities.MethodStatusStatistics, error)
    GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.MethodStatusStatistics, error)
    Create(ctx context.Context, stats *entities.MethodStatusStatistics) error
    Update(ctx context.Context, stats *entities.MethodStatusStatistics) error
}