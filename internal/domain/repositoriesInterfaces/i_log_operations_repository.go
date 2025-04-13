package repositoriesInterfaces

import (
	"context"
	"hub_logging/internal/domain/entities"
)

// ILogOperationsRepository defines the interface for LogOperations persistence.
type ILogOperationsRepository interface {
	Save(ctx context.Context, logOperation entities.LogOperations) error
	FindByLogMessageID(ctx context.Context, logMessageID string) ([]entities.LogOperations, error)
}
