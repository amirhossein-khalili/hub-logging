package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds the configuration values for your application.
type AppConfig struct {
	// Server configuration.
	ServerPort string

	// Database configuration.
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	// Time zone configuration.
	TimeZone string
}

// SetupEnv loads configuration from a .env file and environment variables, then returns an AppConfig instance.
func SetupEnv() (AppConfig, error) {
	// Load .env file if it exists. Ignore error if file is not found.
	_ = godotenv.Load()

	config := AppConfig{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		TimeZone:   os.Getenv("TIME_ZONE"),
	}

	if config.ServerPort == "" {
	config.ServerPort = ":3000"
	}

	// Check if any required configuration is missing.
	if config.DBHost == "" || config.DBUser == "" || config.DBPassword == "" ||
		config.DBName == "" || config.DBPort == "" || config.TimeZone == "" {
		return config, fmt.Errorf("one or more configuration values are missing")
	}

	return config, nil
}
