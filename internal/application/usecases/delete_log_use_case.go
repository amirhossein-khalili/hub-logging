package usecases

import (
	"context"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

type DeleteLogUseCase struct {
	Repo repositoriesInterfaces.ILogMessageRepository
}

func NewDeleteLogUseCase(repo repositoriesInterfaces.ILogMessageRepository) *DeleteLogUseCase {
	return &DeleteLogUseCase{Repo: repo}
}

func (uc *DeleteLogUseCase) Execute(ctx context.Context, id string) error {
	return uc.Repo.Delete(ctx, id)
}
