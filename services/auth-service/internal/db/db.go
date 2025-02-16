package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	DB *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPg(ctx context.Context, connstring string) (*postgres, error) {
	pgOnce.Do(func() {
		DB, err := pgxpool.New(ctx, connstring)
		if err != nil {
			log.Fatal("error getting pool: %w", err)
		}
		pgInstance = &postgres{DB}
	})
	return pgInstance, nil
}

func GetPool() *pgxpool.Pool {
	if pgInstance == nil {
		log.Fatal("Database pool id not initialized")
	}
	return pgInstance.DB
}

func (pg *postgres) Ping(ctx context.Context) error {
	fmt.Println("Ping")
	return pg.DB.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.DB.Close()
}
