package postgres

import (
	"context"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"gorm.io/gorm"
)

// statisticsRepository implements the IStatisticsRepository interface.
type statisticsRepository struct {
	db *gorm.DB
}

// NewStatisticsRepository creates a new instance of statisticsRepository.
func NewStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IStatisticsRepository {
	return &statisticsRepository{db: db}
}

// Save persists the provided Statistics entity in the database.
func (r *statisticsRepository) Save(ctx context.Context, stat entities.Statistics) error {
	return r.db.WithContext(ctx).Create(&stat).Error
}

// FindByRoutePath retrieves Statistics records for a specific route path.
func (r *statisticsRepository) FindByRoutePath(ctx context.Context, routePath string) ([]entities.Statistics, error) {
	var stats []entities.Statistics
	if err := r.db.WithContext(ctx).Where("route_path = ?", routePath).Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

// FindByStatusCode retrieves Statistics records for a specific HTTP status code.
func (r *statisticsRepository) FindByStatusCode(ctx context.Context, statusCode int) ([]entities.Statistics, error) {
	var stats []entities.Statistics
	if err := r.db.WithContext(ctx).Where("status_code = ?", statusCode).Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}
