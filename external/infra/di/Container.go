package di

import (
	"fmt"
	"log"
	"os"
	"time"

	"hub_logging/config"
	"hub_logging/external/infra/models"
	pgRepo "hub_logging/external/infra/repositories/postgres"
	"hub_logging/internal/application/usecases"
	"hub_logging/internal/domain/repositoriesInterfaces"

	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Container holds the dependencies of the application.
type Container struct {
	DB               *gorm.DB
	LogMessageRepo   repositoriesInterfaces.ILogMessageRepository
	CreateLogUseCase *usecases.CreateLogUseCase
}

// NewContainer initializes the DB connection, performs migrations,
// and wires up repositories and use cases.
// It expects the DSN to be available via the "DATABASE_URL" environment variable.
func NewContainer(cfg config.AppConfig) *Container {
	// Build DSN using the DB configuration from cfg.
	// Adjust the DSN format as needed for your Postgres configuration.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  TimeZone=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort  , cfg.TimeZone)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open a GORM DB connection using the Postgres driver.
	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to open DB connection: %v", err)
	}

	// Run auto-migration to create or update the log_message table.
	if err := db.AutoMigrate(&models.LogMessage{}); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	// Create repository and use case instances.
	logRepo := pgRepo.NewLogMessageRepository(db)
	createLogUseCase := usecases.NewCreateLogUseCase(logRepo)

	return &Container{
		DB:               db,
		LogMessageRepo:   logRepo,
		CreateLogUseCase: createLogUseCase,
	}
}
