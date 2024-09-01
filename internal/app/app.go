package app

import (
	"context"
	"delivery-stream-service/config"
	"delivery-stream-service/pkg"
	"log"
	"time"
)

func Run(cfg *config.Config) {
	lg, err := pkg.CreateLogger(cfg.LogFile, "dev")
	if err != nil {
		log.Fatal("Create logger failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
}