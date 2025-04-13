package usecases

import (
	"context"
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

type GetLogsUseCase struct {
	Repo repositoriesInterfaces.ILogMessageRepository
}

func NewGetLogsUseCase(repo repositoriesInterfaces.ILogMessageRepository) *GetLogsUseCase {
	return &GetLogsUseCase{Repo: repo}
}

func (uc *GetLogsUseCase) Execute(ctx context.Context, limit, offset int) ([]entities.LogMessage, error) {
	return uc.Repo.FindWithPagination(ctx, limit, offset) // Pass the context to the repository method
}

func (uc *GetLogsUseCase) ExecuteSingle(ctx context.Context, id string) (entities.LogMessage, error) {
	return uc.Repo.FindByID(ctx, id) // Pass the context to the repository method
}
