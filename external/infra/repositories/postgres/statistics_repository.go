package postgres

import (
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
	"gorm.io/gorm"
)

type statisticsRepository struct {
	db *gorm.DB
}

func NewStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IStatisticsRepository {
	return &statisticsRepository{db: db}
}

func (r *statisticsRepository) Save(stat entities.Statistics) error {
	return r.db.Create(&stat).Error
}

func (r *statisticsRepository) FindByRoutePath(routePath string) ([]entities.Statistics, error) {
	var stats []entities.Statistics
	err := r.db.Where("route_path = ?", routePath).Find(&stats).Error
	return stats, err
}

func (r *statisticsRepository) FindByStatusCode(statusCode int) ([]entities.Statistics, error) {
	var stats []entities.Statistics
	err := r.db.Where("status_code = ?", statusCode).Find(&stats).Error
	return stats, err
}
