package mappers

import (
	"hub_logging/external/infrastructure/models"
	"hub_logging/internal/domain/entities"
)

func ToUserStatisticsModel(entity *entities.UserStatistics) *models.UserStatistics {
	return &models.UserStatistics{
		ID:            entity.ID,
		UserName:      entity.UserName,
		PeriodStart:   entity.PeriodStart,
		PeriodEnd:     entity.PeriodEnd,
		TotalRequests: entity.TotalRequests,
	}
}

func FromUserStatisticsModel(model *models.UserStatistics) *entities.UserStatistics {
	return &entities.UserStatistics{
		ID:            model.ID,
		UserName:      model.UserName,
		PeriodStart:   model.PeriodStart,
		PeriodEnd:     model.PeriodEnd,
		TotalRequests: model.TotalRequests,
	}
}
