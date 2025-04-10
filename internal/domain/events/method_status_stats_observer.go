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

// MethodStatusStatsObserver updates method and status statistics.
type MethodStatusStatsObserver struct {
	methodStatsRepo repositoriesInterfaces.IMethodStatusStatisticsRepository
}

// NewMethodStatusStatsObserver creates a new MethodStatusStatsObserver.
func NewMethodStatusStatsObserver(repo repositoriesInterfaces.IMethodStatusStatisticsRepository) *MethodStatusStatsObserver {
	return &MethodStatusStatsObserver{
		methodStatsRepo: repo,
	}
}

// OnLogCreated processes a log event for method+status statistics.
func (o *MethodStatusStatsObserver) OnLogCreated(logAgg aggregates.LogAggregate) {
	logMsg := logAgg.GetLogMessage()

	periodStart := time.Now().Truncate(24 * time.Hour)
	periodEnd := periodStart.Add(24 * time.Hour)

	existingStats, err := o.methodStatsRepo.GetByPeriod(context.Background(), periodStart, periodEnd)
	if err != nil {
		return
	}

	found := false
	for _, msStat := range existingStats {
		// Compare HTTP method strings; adapt as needed if value objects are used.
		if msStat.HttpMethod.String() == logMsg.HttpMethod {
			msStat.TotalRequests++
			if logMsg.StatusCode >= 200 && logMsg.StatusCode < 300 {
				msStat.SuccessCount++
			} else {
				msStat.ErrorCount++
			}
			_ = o.methodStatsRepo.Update(context.Background(), msStat)
			found = true
			break
		}
	}

	if !found {
		// Create a new MethodStatusStatistics record
		// Wrap the HttpMethod as a value object if required.
		newStat := &entities.MethodStatusStatistics{
			ID:            uuid.New(),
			HttpMethod:    valueobjects.RequestMethod(logMsg.HttpMethod),
			PeriodStart:   periodStart,
			PeriodEnd:     periodEnd,
			TotalRequests: 1,
			SuccessCount: func() int {
				if logMsg.StatusCode >= 200 && logMsg.StatusCode < 300 {
					return 1
				}
				return 0
			}(),
			ErrorCount: func() int {
				if logMsg.StatusCode >= 200 && logMsg.StatusCode < 300 {
					return 0
				}
				return 1
			}(),
		}
		_ = o.methodStatsRepo.Create(context.Background(), newStat)
	}
}
