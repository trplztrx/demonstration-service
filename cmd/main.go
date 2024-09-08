package main

import (
	"delivery-stream-service/config"
	"delivery-stream-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Read config file failed")
	}

	app.Run(cfg)
}