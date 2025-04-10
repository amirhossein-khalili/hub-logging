package events

import (
	"context"
	"time"

	"hub_logging/internal/domain/aggregates"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/google/uuid"
)

// IPStatsObserver updates IP statistics when a new log is created.
type IPStatsObserver struct {
	ipStatsRepo repositoriesInterfaces.IIPStatisticsRepository
}

// NewIPStatsObserver creates a new IPStatsObserver with the provided repository.
func NewIPStatsObserver(repo repositoriesInterfaces.IIPStatisticsRepository) *IPStatsObserver {
	return &IPStatsObserver{
		ipStatsRepo: repo,
	}
}

// OnLogCreated updates or creates IP statistics based on the new log.
func (o *IPStatsObserver) OnLogCreated(logAgg aggregates.LogAggregate) {
	logMsg := logAgg.GetLogMessage()

	periodStart := time.Now().Truncate(24 * time.Hour)
	periodEnd := periodStart.Add(24 * time.Hour)

	// Retrieve existing IP statistics during this period.
	existingStats, err := o.ipStatsRepo.GetByPeriod(context.Background(), periodStart, periodEnd)
	if err != nil {
		// Handle error appropriately (e.g. log the error)
		return
	}

	found := false
	for _, ipStat := range existingStats {
		if ipStat.SourceIP == logMsg.SourceIP {
			ipStat.TotalRequests++
			_ = o.ipStatsRepo.Update(context.Background(), ipStat)
			found = true
			break
		}
	}

	// If not found, create a new statistics record.
	if !found {
		newStats := &entities.IPStatistics{
			ID:            uuid.New(),
			SourceIP:      logMsg.SourceIP,
			PeriodStart:   periodStart,
			PeriodEnd:     periodEnd,
			TotalRequests: 1,
		}
		_ = o.ipStatsRepo.Create(context.Background(), newStats)
	}
}
