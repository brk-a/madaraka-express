package db

import (
    "fmt"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connect() error {
    err := godotenv.Load()
    if err != nil {
        return fmt.Errorf("error loading .env file: %w", err)
    }

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        return fmt.Errorf("DATABASE_URL not set")
    }

    pool, err := pgxpool.New(nil, dbURL)
    if err != nil {
        return fmt.Errorf("unable to create connection pool: %w", err)
    }

    DB = pool
    return nil
}

func Close() {
    if DB != nil {
        DB.Close()
    }
}
