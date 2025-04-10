package events

import (
	"context"
	"time"

	"hub_logging/internal/domain/aggregates"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/google/uuid"
)

// UserStatsObserver updates user-based statistics.
type UserStatsObserver struct {
	userStatsRepo repositoriesInterfaces.IUserStatisticsRepository
}

// NewUserStatsObserver creates a new UserStatsObserver.
func NewUserStatsObserver(repo repositoriesInterfaces.IUserStatisticsRepository) *UserStatsObserver {
	return &UserStatsObserver{
		userStatsRepo: repo,
	}
}

// OnLogCreated updates or creates user statistics based on the new log.
func (o *UserStatsObserver) OnLogCreated(logAgg aggregates.LogAggregate) {
	logMsg := logAgg.GetLogMessage()

	periodStart := time.Now().Truncate(24 * time.Hour)
	periodEnd := periodStart.Add(24 * time.Hour)

	existingStats, err := o.userStatsRepo.GetByPeriod(context.Background(), periodStart, periodEnd)
	if err != nil {
		return
	}

	found := false
	for _, userStat := range existingStats {
		if userStat.UserName == logMsg.UserName {
			userStat.TotalRequests++
			_ = o.userStatsRepo.Update(context.Background(), userStat)
			found = true
			break
		}
	}

	if !found {
		newUserStat := &entities.UserStatistics{
			ID:            uuid.New(),
			UserName:      logMsg.UserName,
			PeriodStart:   periodStart,
			PeriodEnd:     periodEnd,
			TotalRequests: 1,
		}
		_ = o.userStatsRepo.Create(context.Background(), newUserStat)
	}
}
