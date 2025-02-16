package repository

import (
	"auth-services/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UserExist(ctx context.Context, pool *pgxpool.Pool, user models.User) (bool, error) {
	var userExist bool

	err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT * FROM users WHERE email=$1 OR username=$2)`, user.Email, user.Username).Scan(&userExist)
	if err != nil {
		return true, fmt.Errorf("failed to check is user exists: %v", err)
	}

	if userExist {
		return true, fmt.Errorf("user with email %s or username %s already exists", user.Email, user.Username)
	}
	return false, nil
}

func InsertUser(ctx context.Context, pool *pgxpool.Pool, user models.User) error {

	query := `INSERT INTO users 
	(username,
	email,
	password_hash,
	full_name,
	bio,
	profile_picture,
	created_at,
	updated_at,
	is_active,
	is_verified)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	_, err := pool.Exec(ctx, query, user.Username, user.Email, user.PasswordHash, user.FullName, user.Bio, user.ProfilePicture, user.CreatedAt, user.UpdatedAt, user.IsActive, user.IsVerified)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)

	}

	return nil

}
