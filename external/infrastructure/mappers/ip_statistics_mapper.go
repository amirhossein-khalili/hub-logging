package mappers

import (
	"hub_logging/external/infrastructure/models"
	"hub_logging/internal/domain/entities"
)

func ToIPStatisticsModel(entity *entities.IPStatistics) *models.IPStatistics {
	return &models.IPStatistics{
		ID:            entity.ID,
		SourceIP:      entity.SourceIP,
		PeriodStart:   entity.PeriodStart,
		PeriodEnd:     entity.PeriodEnd,
		TotalRequests: entity.TotalRequests,
	}
}

func FromIPStatisticsModel(model *models.IPStatistics) *entities.IPStatistics {
	return &entities.IPStatistics{
		ID:            model.ID,
		SourceIP:      model.SourceIP,
		PeriodStart:   model.PeriodStart,
		PeriodEnd:     model.PeriodEnd,
		TotalRequests: model.TotalRequests,
	}
}
