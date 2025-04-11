package usecases

import (
	"context"
	"time"

	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

// GetIPStatisticsUseCase retrieves IP statistics for a given time period.
type GetMethodStatisticsUseCase struct {
	Repo repositoriesInterfaces.IMethodStatusStatisticsRepository
}

// NewGetIPStatisticsUseCase creates a new instance.
func NewGetMethodStatisticsUseCase(repo repositoriesInterfaces.IMethodStatusStatisticsRepository) *GetMethodStatisticsUseCase {
	return &GetMethodStatisticsUseCase{
		Repo: repo,
	}
}

// Execute returns the IP statistics between the specified start and end times.
func (uc *GetMethodStatisticsUseCase) Execute(ctx context.Context, start, end time.Time) ([]*entities.MethodStatusStatistics, error) {
	return uc.Repo.GetByPeriod(ctx, start, end)
}
