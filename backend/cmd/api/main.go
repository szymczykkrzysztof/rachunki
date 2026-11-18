package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/szymczykkrzysztof/rachunki/backend/internal/httpapi"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := loadDotEnv(); err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("create db pool: %w", err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping db: %w", err)
	}

	r := httpapi.NewRouter(pool)
	return r.Run(":" + port)
}

func loadDotEnv() error {
	for _, path := range []string{".env", "../.env"} {
		err := godotenv.Load(path)
		if err == nil {
			return nil
		}
		if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("load %s: %w", path, err)
		}
	}
	return nil
}
