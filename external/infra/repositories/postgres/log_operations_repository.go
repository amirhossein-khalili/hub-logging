package postgres

import (
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// logOperationsRepository implements the ILogOperationsRepository interface.
type logOperationsRepository struct {
	db *gorm.DB
}

// NewLogOperationsRepository creates a new instance of logOperationsRepository.
func NewLogOperationsRepository(db *gorm.DB) repositoriesInterfaces.ILogOperationsRepository {
	return &logOperationsRepository{db: db}
}

// Save persists the provided LogOperations entity in the database.
func (r *logOperationsRepository) Save(logOperation entities.LogOperations) error {
	return r.db.Create(&logOperation).Error
}

// FindByLogMessageID retrieves all LogOperations associated with a given LogMessage ID.
func (r *logOperationsRepository) FindByLogMessageID(logMessageID string) ([]entities.LogOperations, error) {
	uuidID, err := uuid.Parse(logMessageID)
	if err != nil {
		return nil, err
	}
	var results []entities.LogOperations
	if err := r.db.Where("log_message_id = ?", uuidID).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
