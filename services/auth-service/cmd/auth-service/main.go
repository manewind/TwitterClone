package main

import (
	"auth-services/internal/config"
	"auth-services/internal/db"
	"auth-services/internal/models"
	"auth-services/internal/service"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "aboba")
	})
	pgUser := os.Getenv("PGUSER")
	pgPassword := os.Getenv("PGPASSWORD")
	pgDatabase := os.Getenv("PGDATABASE")
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUser, pgPassword, pgHost, pgPort, pgDatabase)

	pool, err := db.NewPg(context.Background(), connString)
	if err != nil {
		log.Fatal("Error: %w", err)
	}
	defer pool.Close()
	pool.Ping(context.Background())

	sampleUser := models.User{
		Username:       "jsmith25",
		Email:          "john.smith@example.com37",
		FullName:       "John Smith",
		PasswordHash:   "new_hashed_password",
		ProfilePicture: "https://example.com/jsmith.jpg",
		Bio:            "Full-Stack Developer with a passion for open-source projects.",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
		IsActive:       true,
		IsVerified:     false,
	}

	userService := service.NewUserService(pool.DB)

	err = userService.CreateUser(context.Background(), sampleUser)
	if err != nil {
		log.Fatal("error created user: %w", err)
	} else {
		log.Println("User created successfully")
	}

	port := config.NewAppConfig().Port
	log.Printf("Server is starting on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("error starring server: %w\n", err)
	}
}
