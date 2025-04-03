package repositoriesInterfaces

import "hub_logging/internal/domain/entities"

// ILogMessageRepository defines the interface for LogMessage persistence.
type ILogMessageRepository interface {
	// Save persists a single LogMessage.
	Save(logMessage entities.LogMessage) error

	// FindByID retrieves a LogMessage by its unique identifier.
	FindByID(id string) (entities.LogMessage, error)

	// FindAll retrieves all LogMessages.
	FindAll() ([]entities.LogMessage, error)

	// Update modifies an existing LogMessage.
	Update(logMessage entities.LogMessage) error

	// Delete removes a LogMessage by its ID.
	Delete(id string) error
}
