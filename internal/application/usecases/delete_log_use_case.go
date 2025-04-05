package usecases

import "hub_logging/internal/domain/repositoriesInterfaces"

type DeleteLogUseCase struct {
	Repo repositoriesInterfaces.ILogMessageRepository
}

func NewDeleteLogUseCase(repo repositoriesInterfaces.ILogMessageRepository) *DeleteLogUseCase {
	return &DeleteLogUseCase{Repo: repo}
}

func (uc *DeleteLogUseCase) Execute(id string) error {
	return uc.Repo.Delete(id)
}
