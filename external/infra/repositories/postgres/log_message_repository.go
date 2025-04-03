package postgres

import (
	"hub_logging/external/infra/mappers"
	"hub_logging/external/infra/models"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type logMessageRepository struct {
	db *gorm.DB
}

func NewLogMessageRepository(db *gorm.DB) repositoriesInterfaces.ILogMessageRepository {
	return &logMessageRepository{db: db}
}

func (r *logMessageRepository) Save(logMessage entities.LogMessage) error {
	model := mappers.ToModelLogMessage(logMessage)
	return r.db.Create(&model).Error
}

func (r *logMessageRepository) FindByID(id string) (entities.LogMessage, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entities.LogMessage{}, err
	}
	var model models.LogMessage
	err = r.db.First(&model, "id = ?", uuidID).Error
	if err != nil {
		return entities.LogMessage{}, err
	}
	return mappers.ToEntityLogMessage(model), nil
}

func (r *logMessageRepository) FindAll() ([]entities.LogMessage, error) {
	var modelsList []models.LogMessage
	err := r.db.Find(&modelsList).Error
	if err != nil {
		return nil, err
	}

	var entitiesList []entities.LogMessage
	for _, m := range modelsList {
		entitiesList = append(entitiesList, mappers.ToEntityLogMessage(m))
	}
	return entitiesList, nil
}
