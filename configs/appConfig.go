package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	cfg.ServerPort = httpPort
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}
