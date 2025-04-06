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

type MethodStatusStatisticsRepository struct {
	db *gorm.DB
}

func NewMethodStatusStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IMethodStatusStatisticsRepository {
	return &MethodStatusStatisticsRepository{db: db}
}

func (r *MethodStatusStatisticsRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.MethodStatusStatistics, error) {
	var model models.MethodStatusStatistics
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return mappers.FromMethodStatusStatisticsModel(&model), nil
}

func (r *MethodStatusStatisticsRepository) GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.MethodStatusStatistics, error) {
	var modelList []models.MethodStatusStatistics
	if err := r.db.WithContext(ctx).Where("period_start >= ? AND period_end <= ?", start, end).Find(&modelList).Error; err != nil {
		return nil, err
	}
	entityList := make([]*entities.MethodStatusStatistics, len(modelList))
	for i, model := range modelList {
		entityList[i] = mappers.FromMethodStatusStatisticsModel(&model)
	}
	return entityList, nil
}

func (r *MethodStatusStatisticsRepository) Create(ctx context.Context, stats *entities.MethodStatusStatistics) error {
	model := mappers.ToMethodStatusStatisticsModel(stats)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *MethodStatusStatisticsRepository) Update(ctx context.Context, stats *entities.MethodStatusStatistics) error {
	model := mappers.ToMethodStatusStatisticsModel(stats)
	return r.db.WithContext(ctx).Save(model).Error
}
