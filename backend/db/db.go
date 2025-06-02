package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	Pool *pgxpool.Pool
	once sync.Once
)

func Connect() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Printf(".env file not found or failed to load: %v", err)
		}

		connString := os.Getenv("SUPABASE_DATABASE_URL")
		if connString == "" {
			log.Fatal("SUPABASE_DATABASE_URL environment variable not set")
		}

		config, err := pgxpool.ParseConfig(connString)
		if err != nil {
			log.Fatalf("Failed to parse connection string: %v", err)
		}

		config.ConnConfig.RuntimeParams["prepared_statements"] = "false"
		config.ConnConfig.RuntimeParams["statement_cache_mode"] = "none"
		config.ConnConfig.RuntimeParams["application_name"] = "rest_api"
		
		config.MaxConns = 20
		config.MinConns = 2
		
		Pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		conn, err := Pool.Acquire(context.Background())
		if err != nil {
			log.Fatalf("Failed to acquire connection: %v", err)
		}
		defer conn.Release()

		_, err = conn.Exec(context.Background(), "SELECT 1")
		if err != nil {
			log.Fatalf("Failed to execute test query: %v", err)
		}
	})
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}

// DeleteExpiredPayloads deletes all payloads that have expired.
func DeleteExpiredPayloads() error {
	ctx := context.Background()
	_, err := Pool.Exec(ctx, `
		DELETE FROM payloads 
		WHERE (expires_at IS NOT NULL AND expires_at < NOW())
		AND remaining_reads = 0
	`)
	return err
}