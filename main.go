package main

import (
	"hub_logging/configs"
	"hub_logging/internal/api"
	"log"
)

func main() {
	cfg, err := configs.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly: %v \n", err)
	}
	api.StartServer(cfg)
}
