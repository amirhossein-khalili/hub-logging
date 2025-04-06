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

type IPStatisticsRepository struct {
	db *gorm.DB
}

func NewIPStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IIPStatisticsRepository {
	return &IPStatisticsRepository{db: db}
}

func (r *IPStatisticsRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.IPStatistics, error) {
	var model models.IPStatistics
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return mappers.FromIPStatisticsModel(&model), nil
}

func (r *IPStatisticsRepository) GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.IPStatistics, error) {
	var modelList []models.IPStatistics
	if err := r.db.WithContext(ctx).Where("period_start >= ? AND period_end <= ?", start, end).Find(&modelList).Error; err != nil {
		return nil, err
	}
	entityList := make([]*entities.IPStatistics, len(modelList))
	for i, model := range modelList {
		entityList[i] = mappers.FromIPStatisticsModel(&model)
	}
	return entityList, nil
}

func (r *IPStatisticsRepository) Create(ctx context.Context, stats *entities.IPStatistics) error {
	model := mappers.ToIPStatisticsModel(stats)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *IPStatisticsRepository) Update(ctx context.Context, stats *entities.IPStatistics) error {
	model := mappers.ToIPStatisticsModel(stats)
	return r.db.WithContext(ctx).Save(model).Error
}
