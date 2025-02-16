package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Username       string
	Email          string
	FullName       string
	PasswordHash   string
	ProfilePicture string
	Bio            string
	UpdatedAt      time.Time
	CreatedAt      time.Time
	IsActive       bool
	IsVerified     bool
}

type UserService interface {
	UserExist(context.Context, *pgxpool.Pool, User) (bool, error)
	InsertUser(context.Context, *pgxpool.Pool, User) error
	DeleteUser(context.Context, *pgxpool.Pool, User) error
	UpdateUser(context.Context, *pgxpool.Pool, User) error
	ReadUser(context.Context) error
}
