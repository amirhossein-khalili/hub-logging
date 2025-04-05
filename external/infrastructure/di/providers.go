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

// ProvideDB initializes the database connection.
func ProvideDB(cfg config.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.TimeZone)

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
		return nil, fmt.Errorf("failed to open DB connection: %v", err)
	}

	// Run auto-migration to create or update the log_message table.
	if err := db.AutoMigrate(&models.LogMessage{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %v", err)
	}

	return db, nil
}

// ProvideLogMessageRepository initializes the LogMessageRepository.
func ProvideLogMessageRepository(db *gorm.DB) repositoriesInterfaces.ILogMessageRepository {
	return pgRepo.NewLogMessageRepository(db)
}

// ProvideCreateLogUseCase initializes the CreateLogUseCase.
func ProvideCreateLogUseCase(logRepo repositoriesInterfaces.ILogMessageRepository) *usecases.CreateLogUseCase {
	return usecases.NewCreateLogUseCase(logRepo)
}
