package usecases

import (
	"hub_logging/internal/domain/entities"
	"hub_logging/internal/domain/repositoriesInterfaces"
)

type GetLogsUseCase struct {
	Repo repositoriesInterfaces.ILogMessageRepository
}

func NewGetLogsUseCase(repo repositoriesInterfaces.ILogMessageRepository) *GetLogsUseCase {
	return &GetLogsUseCase{Repo: repo}
}

func (uc *GetLogsUseCase) Execute(limit, offset int) ([]entities.LogMessage, error) {
	return uc.Repo.FindWithPagination(limit, offset)
}

func (uc *GetLogsUseCase) ExecuteSingle(id string) (entities.LogMessage, error) {
	return uc.Repo.FindByID(id)
}
