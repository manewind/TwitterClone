package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	Port             string
	ConnectionString string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Port: ":8080",
		ConnectionString: fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGHOST"),
			os.Getenv("PGPORT"),
			os.Getenv("PGDATABASE")),
	}
}
