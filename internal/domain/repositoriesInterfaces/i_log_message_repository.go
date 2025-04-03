package repositoriesInterfaces

import "hub_logging/internal/domain/entities"

// LogMessageRepository defines the interface for LogMessage persistence.
type ILogMessageRepository interface {
	// Save persists a single LogMessage to the underlying storage.
	Save(logMessage entities.LogMessage) error

	// FindByID retrieves a LogMessage by its unique identifier.
	FindByID(id string) (entities.LogMessage, error)

	// FindAll retrieves all LogMessages stored in the system.
	FindAll() ([]entities.LogMessage, error)
}
