package di

import (
	"fmt"
	"log"
	"os"
	"time"

	"hub_logging/config"
	"hub_logging/external/infrastructure/models"
	pgRepo "hub_logging/external/infrastructure/repositories/postgres"
	"hub_logging/internal/application/usecases"
	"hub_logging/internal/domain/repositoriesInterfaces"

	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
--------------------------------------------------------------------
*
*			Container holds the dependencies of the application.
*
--------------------------------------------------------------------
*/
type Container struct {
	DB                         *gorm.DB
	LogMessageRepo             repositoriesInterfaces.ILogMessageRepository
	IPStatisticsRepo           repositoriesInterfaces.IIPStatisticsRepository
	MethodStatusStatisticsRepo repositoriesInterfaces.IMethodStatusStatisticsRepository
	RouteStatisticsRepo        repositoriesInterfaces.IRouteStatisticsRepository
	UserStatisticsRepo         repositoriesInterfaces.IUserStatisticsRepository
	CreateLogUseCase           *usecases.CreateLogUseCase
	GetLogsUseCase             *usecases.GetLogsUseCase
	DeleteLogUseCase           *usecases.DeleteLogUseCase
}

/*
--------------------------------------------------------------------
*
*			InitializeContainer initializes dependencies.
*
--------------------------------------------------------------------
*/
func InitializeContainer(cfg config.AppConfig) (*Container, error) {
	/*--------------------------------------------------------------------
	*			Build DSN using the DB configuration from cfg.
	 --------------------------------------------------------------------*/
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.TimeZone)

	/*--------------------------------------------------------------------
	*						Initialize logger for GORM
	 --------------------------------------------------------------------*/
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	/*--------------------------------------------------------------------
	*						Open a GORM DB connection
	 --------------------------------------------------------------------*/
	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open DB connection: %v", err)
	}

	/*--------------------------------------------------------------------
	*						Run auto-migration
	 --------------------------------------------------------------------*/
	if err := db.AutoMigrate(&models.LogMessage{}, &models.IPStatistics{}, &models.MethodStatusStatistics{}, &models.RouteStatistics{}, &models.UserStatistics{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %v", err)
	}

	/*--------------------------------------------------------------------
	*						INITILIZE REPOSITORIES
	 --------------------------------------------------------------------*/
	logRepo := pgRepo.NewLogMessageRepository(db)
	ipStatsRepo := pgRepo.NewIPStatisticsRepository(db)
	methodStatsRepo := pgRepo.NewMethodStatusStatisticsRepository(db)
	routeStatsRepo := pgRepo.NewRouteStatisticsRepository(db)
	userStatsRepo := pgRepo.NewUserStatisticsRepository(db)

	/*--------------------------------------------------------------------
	*						INITILIZE USECASES
	 --------------------------------------------------------------------*/
	createLogUseCase := usecases.NewCreateLogUseCase(logRepo)
	getLogsUseCase := usecases.NewGetLogsUseCase(logRepo)
	deleteLogUseCase := usecases.NewDeleteLogUseCase(logRepo)

	return &Container{

		DB:                         db,
		LogMessageRepo:             logRepo,
		IPStatisticsRepo:           ipStatsRepo,
		MethodStatusStatisticsRepo: methodStatsRepo,
		RouteStatisticsRepo:        routeStatsRepo,
		UserStatisticsRepo:         userStatsRepo,
		CreateLogUseCase:           createLogUseCase,
		GetLogsUseCase:             getLogsUseCase,
		DeleteLogUseCase:           deleteLogUseCase,
	}, nil
}
