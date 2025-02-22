package main

import (
	"auth-services/internal/config"
	"auth-services/internal/db"
	"auth-services/internal/router"
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	connString := config.NewAppConfig().ConnectionString
	port := config.NewAppConfig().Port

	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to init logger")
	}
	defer logger.Sync()

	godotenv.Load()

	pool, err := db.NewPg(context.Background(), connString)
	if err != nil {
		log.Fatal("Error: %w", err)
	}
	defer pool.Close()
	pool.Ping(context.Background())

	router := router.AuthRouter(pool.DB, logger)

	log.Printf("Server is starting on port %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error starting server: %v\n", err)
	}
}
