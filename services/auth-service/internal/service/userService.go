package service

import (
	"auth-services/internal/models"
	"auth-services/internal/repository"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserService struct {
	userRepo *repository.UserRepository
	logger   *zap.Logger
}

func NewUserService(pool *pgxpool.Pool, logger *zap.Logger) models.UserInterface {
	userRepo := repository.NewUserRepository(pool)
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}

}

func (s *UserService) UserExist(ctx context.Context, user models.User) (bool, error) {
	return s.userRepo.UserExist(ctx, user)
}

func (s *UserService) InsertUser(ctx context.Context, user models.User) error {
	return s.userRepo.InsertUser(ctx, user)
}

func (s *UserService) GetLogger() *zap.Logger {
	return s.logger
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) error {
	exist, err := s.UserExist(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to create suer: %w", err)
	}

	if exist {
		log.Println(user.Email, user.Username)
		return fmt.Errorf("user already exists1212")

	}
	return s.InsertUser(ctx, user)
}
