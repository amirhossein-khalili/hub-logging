package usecases

import (
	"context"
	"time"

	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

// GetIPStatisticsUseCase retrieves IP statistics for a given time period.
type GetIPStatisticsUseCase struct {
	Repo repositoriesInterfaces.IIPStatisticsRepository
}

// NewGetIPStatisticsUseCase creates a new instance.
func NewGetIPStatisticsUseCase(repo repositoriesInterfaces.IIPStatisticsRepository) *GetIPStatisticsUseCase {
	return &GetIPStatisticsUseCase{
		Repo: repo,
	}
}

// Execute returns the IP statistics between the specified start and end times.
func (uc *GetIPStatisticsUseCase) Execute(ctx context.Context, start, end time.Time) ([]*entities.IPStatistics, error) {
	return uc.Repo.GetByPeriod(ctx, start, end)
}
