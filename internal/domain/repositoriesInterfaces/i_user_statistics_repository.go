package repositoriesInterfaces

import (
	"context"
	"hub_logging/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type IUserStatisticsRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.UserStatistics, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.UserStatistics, error)
	Create(ctx context.Context, stats *entities.UserStatistics) error
	Update(ctx context.Context, stats *entities.UserStatistics) error
}
