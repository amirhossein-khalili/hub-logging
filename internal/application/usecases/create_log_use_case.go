package usecases

import (
	"context"
	"time"

	"hub_logging/internal/application/dtos"
	"hub_logging/internal/domain/aggregates"
	"hub_logging/internal/domain/events"
	"hub_logging/internal/domain/repositoriesInterfaces"
	"hub_logging/internal/domain/valueobjects"
)

type CreateLogUseCase struct {
	LogRepo           repositoriesInterfaces.ILogMessageRepository
	LogEventPublisher events.ILogEventPublisher
}

func NewCreateLogUseCase(
	logRepo repositoriesInterfaces.ILogMessageRepository,
	publisher events.ILogEventPublisher,
) *CreateLogUseCase {
	return &CreateLogUseCase{
		LogRepo:           logRepo,
		LogEventPublisher: publisher,
	}
}

func (uc *CreateLogUseCase) Execute(ctx context.Context, input dtos.CreateLogDTO) error {
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

	// Build the LogAggregate using current time
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

	// Persist the new LogMessage using the repository, passing the context
	if err := uc.LogRepo.Save(ctx, aggregate.GetLogMessage()); err != nil {
		return err
	}

	// Publish the event so observers can update statistics
	uc.LogEventPublisher.PublishLogCreated(*aggregate)

	return nil
}
