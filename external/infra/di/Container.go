package di

import (
	"hub_logging/external/infra/repositories/postgres"
	"hub_logging/internal/application/usecases"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"gorm.io/gorm"
)

type Container struct {
	LogMessageRepo   repositoriesInterfaces.ILogMessageRepository
	CreateLogUseCase *usecases.CreateLogUseCase
}

func NewContainer(db *gorm.DB) *Container {
	logRepo := postgres.NewLogMessageRepository(db)
	return &Container{
		LogMessageRepo:   logRepo,
		CreateLogUseCase: usecases.NewCreateLogUseCase(logRepo),
	}
}
