package models

import "time"

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
	CreateUser() error
}
