package mappers

import (
	"hub_logging/external/infrastructure/models"
	"hub_logging/internal/domain/entities"
)

func ToRouteStatisticsModel(entity *entities.RouteStatistics) *models.RouteStatistics {
	return &models.RouteStatistics{
		ID:            entity.ID,
		RoutePath:     entity.RoutePath,
		PeriodStart:   entity.PeriodStart,
		PeriodEnd:     entity.PeriodEnd,
		TotalRequests: entity.TotalRequests,
		SuccessCount:  entity.SuccessCount,
		ErrorCount:    entity.ErrorCount,
	}
}

func FromRouteStatisticsModel(model *models.RouteStatistics) *entities.RouteStatistics {
	return &entities.RouteStatistics{
		ID:            model.ID,
		RoutePath:     model.RoutePath,
		PeriodStart:   model.PeriodStart,
		PeriodEnd:     model.PeriodEnd,
		TotalRequests: model.TotalRequests,
		SuccessCount:  model.SuccessCount,
		ErrorCount:    model.ErrorCount,
	}
}
