package usecases

import (
	"context"
	"time"

	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

// GetIPStatisticsUseCase retrieves IP statistics for a given time period.
type GetUserStatisticsUseCase struct {
	Repo repositoriesInterfaces.IUserStatisticsRepository
}

// NewGetIPStatisticsUseCase creates a new instance.
func NewGetUserStatisticsUseCase(repo repositoriesInterfaces.IUserStatisticsRepository) *GetUserStatisticsUseCase {
	return &GetUserStatisticsUseCase{
		Repo: repo,
	}
}

// Execute returns the IP statistics between the specified start and end times.
func (uc *GetUserStatisticsUseCase) Execute(ctx context.Context, start, end time.Time) ([]*entities.UserStatistics, error) {
	return uc.Repo.GetByPeriod(ctx, start, end)
}
