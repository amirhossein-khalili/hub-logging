package mappers

import (
	"hub_logging/external/infrastructure/models"
	"hub_logging/internal/domain/entities"
)

func ToMethodStatusStatisticsModel(entity *entities.MethodStatusStatistics) *models.MethodStatusStatistics {
	return &models.MethodStatusStatistics{
		ID:            entity.ID,
		HttpMethod:    entity.HttpMethod,
		PeriodStart:   entity.PeriodStart,
		PeriodEnd:     entity.PeriodEnd,
		TotalRequests: entity.TotalRequests,
		SuccessCount:  entity.SuccessCount,
		ErrorCount:    entity.ErrorCount,
	}
}

func FromMethodStatusStatisticsModel(model *models.MethodStatusStatistics) *entities.MethodStatusStatistics {
	return &entities.MethodStatusStatistics{
		ID:            model.ID,
		HttpMethod:    model.HttpMethod,
		PeriodStart:   model.PeriodStart,
		PeriodEnd:     model.PeriodEnd,
		TotalRequests: model.TotalRequests,
		SuccessCount:  model.SuccessCount,
		ErrorCount:    model.ErrorCount,
	}
}
