package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/egeuysall/rest/db"
	"github.com/go-chi/chi/v5"
)

type Payload struct {
	Data           json.RawMessage `json:"data"`
	TTL            int            `json:"ttl,omitempty"`
	ExpiresIn      int            `json:"expires_in,omitempty"`
	RemainingReads int            `json:"remaining_reads,omitempty"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, please go to /v1/payload to POST."))
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running."))
}

func StorePayload(data json.RawMessage, expire int, times int) error {
	var expiresAt *time.Time
	if expire != -1 {
		t := time.Now().Add(time.Duration(expire) * time.Minute)
		expiresAt = &t
	}

	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	var id string
	var query string
	if expiresAt != nil {
		query = fmt.Sprintf(
			"INSERT INTO payloads (data, expires_at, remaining_reads) VALUES ('%s', '%s', %d) RETURNING id",
			data, expiresAt.Format(time.RFC3339), times,
		)
	} else {
		query = fmt.Sprintf(
			"INSERT INTO payloads (data, expires_at, remaining_reads) VALUES ('%s', NULL, %d) RETURNING id",
			data, times,
		)
	}
	err = tx.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to store payload: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	
	return nil
}

func createPayloadHandler(w http.ResponseWriter, r *http.Request) {
	rawData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var payload Payload
	if err := json.Unmarshal(rawData, &payload); err != nil {
		http.Error(w, "Invalid payload format", http.StatusBadRequest)
		return
	}

	if !json.Valid(payload.Data) {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var expiresAt *time.Time
	if payload.ExpiresIn > 0 {
		t := time.Now().Add(time.Duration(payload.ExpiresIn) * time.Minute)
		expiresAt = &t
	} else if payload.TTL > 0 {
		t := time.Now().Add(time.Duration(payload.TTL) * time.Minute)
		expiresAt = &t
	} else if payload.ExpiresIn != -1 && payload.TTL != -1 {
		t := time.Now().Add(10 * time.Minute)
		expiresAt = &t
	}

	if payload.RemainingReads == 0 {
		payload.RemainingReads = 1
	}

	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	var id string
	var query string
	if expiresAt != nil {
		query = fmt.Sprintf(
			"INSERT INTO payloads (data, expires_at, remaining_reads) VALUES ('%s', '%s', %d) RETURNING id",
			payload.Data, expiresAt.Format(time.RFC3339), payload.RemainingReads,
		)
	} else {
		query = fmt.Sprintf(
			"INSERT INTO payloads (data, expires_at, remaining_reads) VALUES ('%s', NULL, %d) RETURNING id",
			payload.Data, payload.RemainingReads,
		)
	}
	err = tx.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		http.Error(w, "Failed to store payload", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id": id,
		"url": fmt.Sprintf("https://rest.egeuysal.com/payload/%s", id),
	})
}

func getPayloadHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	var data json.RawMessage
	var remainingReads int
	var expiresAt *time.Time

	query := fmt.Sprintf(
		"SELECT data, remaining_reads, expires_at FROM payloads WHERE id = '%s' FOR UPDATE",
		id,
	)
	err = tx.QueryRow(ctx, query).Scan(&data, &remainingReads, &expiresAt)
	if err != nil {
		http.Error(w, "Payload not found", http.StatusNotFound)
		return
	}

	if expiresAt != nil && expiresAt.Before(time.Now()) {
		query = fmt.Sprintf("DELETE FROM payloads WHERE id = '%s'", id)
		_, err = tx.Exec(ctx, query)
		if err != nil {
			http.Error(w, "Failed to delete expired payload", http.StatusInternalServerError)
			return
		}
		if err := tx.Commit(ctx); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Payload has expired", http.StatusGone)
		return
	}

	if remainingReads != -1 {
		if remainingReads <= 0 {
			http.Error(w, "No remaining reads", http.StatusGone)
			return
		}

		remainingReads--
		if remainingReads == 0 {
			query = fmt.Sprintf(
				"UPDATE payloads SET remaining_reads = 0 WHERE id = '%s'",
				id,
			)
			_, err = tx.Exec(ctx, query)
			if err != nil {
				http.Error(w, "Failed to update remaining reads", http.StatusInternalServerError)
				return
			}

			if err := tx.Commit(ctx); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"data": data,
				"remaining_reads": 0,
				"expires_at": expiresAt,
			})
			return
		}

		query = fmt.Sprintf(
			"UPDATE payloads SET remaining_reads = %d WHERE id = '%s'",
			remainingReads, id,
		)
		_, err = tx.Exec(ctx, query)
		if err != nil {
			http.Error(w, "Failed to update remaining reads", http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(ctx); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": data,
		"remaining_reads": remainingReads,
		"expires_at": expiresAt,
	})
}

func deletePayloadHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	ctx := context.Background()
	result, err := db.Pool.Exec(ctx, "DELETE FROM payloads WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete payload", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected() == 0 {
		http.Error(w, "Payload not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}