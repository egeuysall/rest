package db

import (
    "testing"
	"context"
)

func TestConnectAndClose(t *testing.T) {
    Connect()
    if Conn == nil {
        t.Fatal("Connection is nil")
    }

    if err := Conn.Ping(context.Background()); err != nil {
        t.Fatalf("Ping failed: %v", err)
    }

    Close()
}