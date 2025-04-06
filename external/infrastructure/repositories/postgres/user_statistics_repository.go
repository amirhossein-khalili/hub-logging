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

type UserStatisticsRepository struct {
	db *gorm.DB
}

func NewUserStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IUserStatisticsRepository {
	return &UserStatisticsRepository{db: db}
}

func (r *UserStatisticsRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.UserStatistics, error) {
	var model models.UserStatistics
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return mappers.FromUserStatisticsModel(&model), nil
}

func (r *UserStatisticsRepository) GetByPeriod(ctx context.Context, start, end time.Time) ([]*entities.UserStatistics, error) {
	var modelList []models.UserStatistics
	if err := r.db.WithContext(ctx).Where("period_start >= ? AND period_end <= ?", start, end).Find(&modelList).Error; err != nil {
		return nil, err
	}
	entityList := make([]*entities.UserStatistics, len(modelList))
	for i, model := range modelList {
		entityList[i] = mappers.FromUserStatisticsModel(&model)
	}
	return entityList, nil
}

func (r *UserStatisticsRepository) Create(ctx context.Context, stats *entities.UserStatistics) error {
	model := mappers.ToUserStatisticsModel(stats)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *UserStatisticsRepository) Update(ctx context.Context, stats *entities.UserStatistics) error {
	model := mappers.ToUserStatisticsModel(stats)
	return r.db.WithContext(ctx).Save(model).Error
}
