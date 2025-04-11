package usecases

import (
	"context"
	"time"

	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

// GetRouteStatisticsUseCase retrieves route statistics for a given period.
type GetRouteStatisticsUseCase struct {
	Repo repositoriesInterfaces.IRouteStatisticsRepository
}

// NewGetRouteStatisticsUseCase creates a new instance.
func NewGetRouteStatisticsUseCase(repo repositoriesInterfaces.IRouteStatisticsRepository) *GetRouteStatisticsUseCase {
	return &GetRouteStatisticsUseCase{
		Repo: repo,
	}
}

// Execute returns the Route statistics for the given period.
func (uc *GetRouteStatisticsUseCase) Execute(ctx context.Context, start, end time.Time) ([]*entities.RouteStatistics, error) {
	return uc.Repo.GetByPeriod(ctx, start, end)
}
