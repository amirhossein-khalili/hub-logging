package postgres

import (
	"context"
	"hub_logging/external/infrastructure/mappers"
	"hub_logging/external/infrastructure/models"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RouteStatisticsRepository struct {
	db *gorm.DB
}

func NewRouteStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IRouteStatisticsRepository {
	return &RouteStatisticsRepository{db: db}
}

func (r *RouteStatisticsRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.RouteStatistics, error) {
	var model models.RouteStatistics
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return mappers.FromRouteStatisticsModel(&model), nil
}

func (r *RouteStatisticsRepository) GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.RouteStatistics, error) {
	var modelList []models.RouteStatistics
	if err := r.db.WithContext(ctx).Where("period_start >= ? AND period_end <= ?", start, end).Find(&modelList).Error; err != nil {
		return nil, err
	}
	entityList := make([]*entities.RouteStatistics, len(modelList))
	for i, model := range modelList {
		entityList[i] = mappers.FromRouteStatisticsModel(&model)
	}
	return entityList, nil
}

func (r *RouteStatisticsRepository) Create(ctx context.Context, stats *entities.RouteStatistics) error {
	model := mappers.ToRouteStatisticsModel(stats)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *RouteStatisticsRepository) Update(ctx context.Context, stats *entities.RouteStatistics) error {
	model := mappers.ToRouteStatisticsModel(stats)
	return r.db.WithContext(ctx).Save(model).Error
}
