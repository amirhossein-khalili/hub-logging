package usecases

import (
	"hub_logging/internal/application/dtos"
	"hub_logging/internal/domain/aggregates"
	"hub_logging/internal/domain/repositoriesInterfaces"
	"hub_logging/internal/domain/valueobjects"
	"time"
)

type CreateLogUseCase struct {
	LogRepo repositoriesInterfaces.ILogMessageRepository
}

func NewCreateLogUseCase(logRepo repositoriesInterfaces.ILogMessageRepository) *CreateLogUseCase {
	return &CreateLogUseCase{LogRepo: logRepo}
}

func (uc *CreateLogUseCase) Execute(input dtos.CreateLogDTO) error {
	statusCode, err := valueobjects.NewStatusCode(input.StatusCode)
	if err != nil {
		return err
	}
	method, err := valueobjects.NewRequestMethod(input.HttpMethod)
	if err != nil {
		return err
	}
	route, err := valueobjects.NewRoutePath(input.RoutePath)
	if err != nil {
		return err
	}

	aggregate, err := aggregates.NewLogAggregate(
		time.Now(),
		statusCode,
		method,
		route,
		input.Message,
		input.UserName,
		input.DestHostname,
		input.SourceIP,
	)
	if err != nil {
		return err
	}

	return uc.LogRepo.Save(aggregate.GetLogMessage())
}
