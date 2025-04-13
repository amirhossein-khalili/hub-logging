package repositoriesInterfaces

import (
	"context"
	"hub_logging/internal/domain/entities"
)

// ILogMessageRepository defines the interface for LogMessage persistence.
type ILogMessageRepository interface {
	Save(ctx context.Context, logMessage entities.LogMessage) error
	FindByID(ctx context.Context, id string) (entities.LogMessage, error)
	FindAll(ctx context.Context) ([]entities.LogMessage, error)
	FindWithPagination(ctx context.Context, limit, offset int) ([]entities.LogMessage, error)
	Update(ctx context.Context, logMessage entities.LogMessage) error
	Delete(ctx context.Context, id string) error
}
