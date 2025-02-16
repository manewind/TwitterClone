package service

import (
	"auth-services/internal/models"
	"auth-services/internal/repository"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	pool *pgxpool.Pool
	// userService models.UserService
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{
		pool: pool,
		// userService: userService ,
	}

}

func (s *UserService) CreateUser(ctx context.Context, user models.User) error {
	exist, err := repository.UserExist(ctx, s.pool, user)

	if err != nil {
		return fmt.Errorf("failed to create suer: %w", err)
	}

	if exist {
		log.Println(user.Email, user.Username)
		return fmt.Errorf("user already exists1212")

	}
	return repository.InsertUser(ctx, s.pool, user)
}
