package main

import (
	"auth-services/internal/config"
	"auth-services/internal/db"
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "aboba")
	})

	pool, err := db.NewPg(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error: %w", err)
	}
	defer pool.Close()
	pool.Ping(context.Background())

	port := config.NewAppConfig().Port
	log.Printf("Server is starting on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("error starring server: %w\n", err)
	}
}
