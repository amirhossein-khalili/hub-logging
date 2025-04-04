package main

import (
	"hub_logging/config"
	"hub_logging/external/presentation/api/server"
	"log"
)

func main() { // Load application configuration.
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	server.StartServer(cfg)
}
