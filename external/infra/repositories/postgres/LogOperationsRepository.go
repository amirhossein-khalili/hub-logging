package postgres

import (
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type logOperationsRepository struct {
	db *gorm.DB
}

func NewLogOperationsRepository(db *gorm.DB) repositoriesInterfaces.ILogOperationsRepository {
	return &logOperationsRepository{db: db}
}

func (r *logOperationsRepository) Save(logOperation entities.LogOperations) error {
	return r.db.Create(&logOperation).Error
}

func (r *logOperationsRepository) FindByLogMessageID(logMessageID string) ([]entities.LogOperations, error) {
	uuidID, err := uuid.Parse(logMessageID)
	if err != nil {
		return nil, err
	}
	var results []entities.LogOperations
	err = r.db.Where("log_message_id = ?", uuidID).Find(&results).Error
	return results, err
}
