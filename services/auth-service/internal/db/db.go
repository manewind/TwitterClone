package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPg(ctx context.Context, connstring string) (*postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connstring)
		if err != nil {
			log.Fatal("error getting pool: %w", err)
		}
		pgInstance = &postgres{db}
	})
	return pgInstance, nil
}

func (pg *postgres) Ping(ctx context.Context) error {
	fmt.Println("Ping")
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}
