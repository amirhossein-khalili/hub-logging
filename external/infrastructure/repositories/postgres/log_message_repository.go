package postgres

import (
	"hub_logging/external/infrastructure/mappers"
	"hub_logging/external/infrastructure/models"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// logMessageRepository implements the ILogMessageRepository interface.
type logMessageRepository struct {
	db *gorm.DB
}

// NewLogMessageRepository creates a new instance of logMessageRepository.
func NewLogMessageRepository(db *gorm.DB) repositoriesInterfaces.ILogMessageRepository {
	return &logMessageRepository{db: db}
}

// Save persists the provided LogMessage entity in the database.
func (r *logMessageRepository) Save(logMessage entities.LogMessage) error {
	model := mappers.ToModelLogMessage(logMessage)
	return r.db.Create(&model).Error
}

// FindByID retrieves a LogMessage entity by its ID.
func (r *logMessageRepository) FindByID(id string) (entities.LogMessage, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entities.LogMessage{}, err
	}
	var model models.LogMessage
	if err := r.db.First(&model, "id = ?", uuidID).Error; err != nil {
		return entities.LogMessage{}, err
	}
	return mappers.ToEntityLogMessage(model), nil
}

// FindAll retrieves all LogMessage entities from the database.
func (r *logMessageRepository) FindAll() ([]entities.LogMessage, error) {
	var modelsList []models.LogMessage
	if err := r.db.Find(&modelsList).Error; err != nil {
		return nil, err
	}
	entitiesList := make([]entities.LogMessage, len(modelsList))
	for i, m := range modelsList {
		entitiesList[i] = mappers.ToEntityLogMessage(m)
	}
	return entitiesList, nil
}

// FindWithPagination retrieves logs with pagination.
func (r *logMessageRepository) FindWithPagination(limit, offset int) ([]entities.LogMessage, error) {
	var modelsList []models.LogMessage
	err := r.db.Limit(limit).Offset(offset).Find(&modelsList).Error
	if err != nil {
		return nil, err
	}

	// Convert models to entities.
	var entitiesList []entities.LogMessage
	for _, m := range modelsList {
		entitiesList = append(entitiesList, mappers.ToEntityLogMessage(m))
	}
	return entitiesList, nil
}

// Update modifies an existing LogMessage in the database.
func (r *logMessageRepository) Update(logMessage entities.LogMessage) error {
	// GORM's Save will update all fields.
	return r.db.Save(&logMessage).Error
}

// Delete removes a LogMessage from the database by its ID.
func (r *logMessageRepository) Delete(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&models.LogMessage{}, "id = ?", uuidID).Error
}
