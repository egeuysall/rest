package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var (
    Conn *pgx.Conn
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

        var err error
        Conn, err = pgx.Connect(context.Background(), connString)
        if err != nil {
            log.Fatalf("Failed to connect to Supabase Postgres: %v", err)
        }

        log.Println("Connected to Supabase Postgres.")
    })
}

func Close() {
    if Conn != nil {
        Conn.Close(context.Background())
        log.Println("Closed database connection.")
    }
}