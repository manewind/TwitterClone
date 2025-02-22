package models

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type User struct {
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	PasswordHash   string    `json:"password_hash"`
	ProfilePicture string    `json:"profile_picture"`
	Bio            string    `json:"bio"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	IsActive       bool      `json:"is_active"`
	IsVerified     bool      `json:"is_verified"`
}

type UserInterface interface {
	UserExist(context.Context, User) (bool, error)
	InsertUser(context.Context, User) error
	CreateUser(context.Context, User) error
	GetLogger() *zap.Logger
	// DeleteUser(context.Context, User) error
	// UpdateUser(context.Context, User) error
	// ReadUser(context.Context) error
}
