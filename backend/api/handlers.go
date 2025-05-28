package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/egeuysall/rest/db"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Payload struct {
	Data           json.RawMessage `json:"data"`
	TTL            int             `json:"ttl"`
	RemainingReads int             `json:"reads"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, please go to /v1/payload to POST."))
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running."))
}

func createPayloadHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	expiresAt := time.Now().Add(time.Duration(p.TTL) * time.Second)

	_, err := db.Conn.Exec(context.Background(), `
		INSERT INTO payloads (id, data, expires_at, remaining_reads)
		VALUES ($1, $2, $3, $4)
	`, id, p.Data, expiresAt, p.RemainingReads)
	if err != nil {
		http.Error(w, "Failed to store payload", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func getPayloadHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Requested payload ID: " + id))
}

func deletePayloadHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Requested payload ID: " + id))
}