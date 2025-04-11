package di

import (
	pgRepo "hub_logging/external/infrastructure/repositories/postgres"
	"hub_logging/internal/application/usecases"
	"hub_logging/internal/domain/events"
	"hub_logging/internal/domain/repositoriesInterfaces"

	"gorm.io/gorm"
)

// ProvideLogMessageRepository initializes the LogMessageRepository.
func ProvideLogMessageRepository(db *gorm.DB) repositoriesInterfaces.ILogMessageRepository {
	return pgRepo.NewLogMessageRepository(db)
}

// ProvideIPStatisticsRepository initializes the IPStatisticsRepository.
func ProvideIPStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IIPStatisticsRepository {
	return pgRepo.NewIPStatisticsRepository(db)
}

// ProvideMethodStatusStatisticsRepository initializes the MethodStatusStatisticsRepository.
func ProvideMethodStatusStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IMethodStatusStatisticsRepository {
	return pgRepo.NewMethodStatusStatisticsRepository(db)
}

// ProvideRouteStatisticsRepository initializes the RouteStatisticsRepository.
func ProvideRouteStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IRouteStatisticsRepository {
	return pgRepo.NewRouteStatisticsRepository(db)
}

// ProvideUserStatisticsRepository initializes the UserStatisticsRepository.
func ProvideUserStatisticsRepository(db *gorm.DB) repositoriesInterfaces.IUserStatisticsRepository {
	return pgRepo.NewUserStatisticsRepository(db)
}

// ProvideCreateLogUseCase initializes the CreateLogUseCase.
func ProvideCreateLogUseCase(logRepo repositoriesInterfaces.ILogMessageRepository, publisher events.ILogEventPublisher) *usecases.CreateLogUseCase {
	return usecases.NewCreateLogUseCase(logRepo, publisher)
}

// ProvideGetRouteStatisticsUseCase initializes the GetRouteStatisticsUseCase.
func ProvideGetRouteStatisticsUseCase(repo repositoriesInterfaces.IRouteStatisticsRepository) *usecases.GetRouteStatisticsUseCase {
	return usecases.NewGetRouteStatisticsUseCase(repo)
}

// ProvideGetIPStatisticsUseCase initializes the GetIPStatisticsUseCase.
func ProvideGetIPStatisticsUseCase(repo repositoriesInterfaces.IIPStatisticsRepository) *usecases.GetIPStatisticsUseCase {
	return usecases.NewGetIPStatisticsUseCase(repo)
}

// ProvideGetMethodStatisticsUseCase initialize the GetMethodStatisticsUseCase
func ProvideGetMethodStatisticsUseCase(repo repositoriesInterfaces.IMethodStatusStatisticsRepository) *usecases.GetMethodStatisticsUseCase {
	return usecases.NewGetMethodStatisticsUseCase(repo)
}

// ProvideGetUserStatisticsUseCase initialize the GetUserStatisticsUseCase
func ProvideGetUserStatisticsUseCase(repo repositoriesInterfaces.IUserStatisticsRepository) *usecases.GetUserStatisticsUseCase {
	return usecases.NewGetUserStatisticsUseCase(repo)
}
