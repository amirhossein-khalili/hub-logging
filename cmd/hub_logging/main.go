package main

import (
	"hub_logging/config"
	"hub_logging/external/presentation/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly: %v \n", err)
	}
	api.StartServer(cfg)
}
