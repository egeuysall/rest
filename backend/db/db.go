package db

import (
    "context"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() {
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
}

func Close() {
    if Conn != nil {
        Conn.Close(context.Background())
    }
}