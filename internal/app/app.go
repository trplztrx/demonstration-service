package app

import (
	"context"
	"delivery-stream-service/config"
	"delivery-stream-service/pkg"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	lg, err := pkg.CreateLogger(cfg.LogFile, "dev")
	if err != nil {
		log.Fatal("Create logger failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password ,cfg.Host, cfg.Port, cfg.Db.Db)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Failed connection to postgres: %v", err.Error())
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		lg.Warn("Failed to ping database", zap.Error(err))
	} else {
		lg.Info("Connected to PostgreSQL database")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.With(mdware.ipWhitelistMiddleware).Get("/admin/generate", handler.createCoordinatorOrderHandler)
	// router.Get("/orders/{id}", )
	// router.Get("/orders", )
	router.Get("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sosi!"))
	})
	
	lg.Info("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		lg.Fatal("Failed to start server", zap.Error(err))
	}
}