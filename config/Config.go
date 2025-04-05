package config

import (
	"fmt"

	"github.com/spf13/viper"
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

// SetupEnv loads configuration from different sources (e.g., .env, environment variables, etc.)
func SetupEnv() (AppConfig, error) {
	// Set the config file name and path.
	viper.SetConfigName(".env") // .env file as the config file
	viper.SetConfigType("env")  // File format
	viper.AddConfigPath(".")    // Search for the .env file in the current directory
	viper.AutomaticEnv()        // Automatically read from environment variables

	// Read from the config file, if it exists.
	if err := viper.ReadInConfig(); err != nil {
		// If the .env file does not exist, it will ignore this error.
		// You can log it if necessary, but it’s not critical to the functionality.
		fmt.Println("No .env file found, continuing with environment variables")
	}

	// Set default values for configuration fields (optional).
	viper.SetDefault("SERVER_PORT", ":3000")

	// Load configuration from environment variables using Viper.
	config := AppConfig{
		ServerPort: viper.GetString("SERVER_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetString("DB_PORT"),
		TimeZone:   viper.GetString("TIME_ZONE"),
	}

	// Validate that the configuration values are set.
	if config.DBHost == "" || config.DBUser == "" || config.DBPassword == "" ||
		config.DBName == "" || config.DBPort == "" || config.TimeZone == "" {
		return config, fmt.Errorf("one or more configuration values are missing")
	}

	return config, nil
}
