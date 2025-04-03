package postgres

import (
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
func (r *statisticsRepository) Save(stat entities.Statistics) error {
	return r.db.Create(&stat).Error
}

// FindByRoutePath retrieves Statistics records for a specific route path.
func (r *statisticsRepository) FindByRoutePath(routePath string) ([]entities.Statistics, error) {
	var stats []entities.Statistics
	if err := r.db.Where("route_path = ?", routePath).Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}

// FindByStatusCode retrieves Statistics records for a specific HTTP status code.
func (r *statisticsRepository) FindByStatusCode(statusCode int) ([]entities.Statistics, error) {
	var stats []entities.Statistics
	if err := r.db.Where("status_code = ?", statusCode).Find(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}
