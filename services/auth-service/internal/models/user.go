package models

import (
	"context"
	"time"
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

type UserInterface interface {
	UserExist(context.Context, User) (bool, error)
	InsertUser(context.Context, User) error
	CreateUser(context.Context, User) error
	// DeleteUser(context.Context, User) error
	// UpdateUser(context.Context, User) error
	// ReadUser(context.Context) error
}
