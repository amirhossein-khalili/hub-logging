package events

import (
	"context"
	"time"

	"hub_logging/internal/domain/aggregates"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
	"hub_logging/internal/domain/valueobjects"

	"github.com/google/uuid"
)

// RouteStatsObserver updates route statistics based on new logs.
type RouteStatsObserver struct {
	routeStatsRepo repositoriesInterfaces.IRouteStatisticsRepository
}

// NewRouteStatsObserver creates a new RouteStatsObserver.
func NewRouteStatsObserver(repo repositoriesInterfaces.IRouteStatisticsRepository) *RouteStatsObserver {
	return &RouteStatsObserver{
		routeStatsRepo: repo,
	}
}

// OnLogCreated processes a new log event for route statistics update.
func (o *RouteStatsObserver) OnLogCreated(logAgg aggregates.LogAggregate) {
	logMsg := logAgg.GetLogMessage()

	periodStart := time.Now().Truncate(24 * time.Hour)
	periodEnd := periodStart.Add(24 * time.Hour)

	existingStats, err := o.routeStatsRepo.GetByPeriod(context.Background(), periodStart, periodEnd)
	if err != nil {
		return
	}

	found := false
	for _, routeStat := range existingStats {
		// Compare using the string representation of the RoutePath
		if routeStat.RoutePath.String() == logMsg.RoutePath {
			routeStat.TotalRequests++
			if logMsg.StatusCode >= 200 && logMsg.StatusCode < 300 {
				routeStat.SuccessCount++
			} else {
				routeStat.ErrorCount++
			}
			_ = o.routeStatsRepo.Update(context.Background(), routeStat)
			found = true
			break
		}
	}

	if !found {
		newStat := &entities.RouteStatistics{
			ID: uuid.New(),
			// Convert the string value into a valueobjects.RoutePath type.
			RoutePath:     valueobjects.RoutePath(logMsg.RoutePath),
			PeriodStart:   periodStart,
			PeriodEnd:     periodEnd,
			TotalRequests: 1,
			SuccessCount: func() int {
				if logMsg.StatusCode >= 200 && logMsg.StatusCode < 300 {
					return 1
				} else {
					return 0
				}
			}(),
			ErrorCount: func() int {
				if logMsg.StatusCode >= 200 && logMsg.StatusCode < 300 {
					return 0
				} else {
					return 1
				}
			}(),
		}
		_ = o.routeStatsRepo.Create(context.Background(), newStat)
	}
}
