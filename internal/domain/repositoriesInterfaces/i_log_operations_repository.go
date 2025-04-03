package repositoriesInterfaces

import "hub_logging/internal/domain/entities"

// LogOperationsRepository defines the interface for LogOperations persistence.
type ILogOperationsRepository interface {
	// Save persists a single LogOperations entity.
	Save(logOperation entities.LogOperations) error

	// FindByLogMessageID retrieves all LogOperations associated with a specific LogMessage ID.
	FindByLogMessageID(logMessageID string) ([]entities.LogOperations, error)
}
