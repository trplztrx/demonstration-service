package main

import (
	"delivery-stream-service/config"
	"delivery-stream-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Failed read config file")
	}

	app.Run(cfg)
}