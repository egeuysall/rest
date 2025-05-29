package api

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		apiKey := os.Getenv("REST_API_KEY")
		if apiKey == "" {
			log.Fatal("REST_API_KEY environment variable not set")
		}

		if parts[1] != apiKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Use(ConfigureCORS())
	r.Use(NewRateLimiter(1, 3).EnforceRateLimit())

	r.Get("/", handleRoot)
	r.Get("/health", checkHealth)

	r.Route("/v1", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/payload", createPayloadHandler)
		r.Get("/payload/{id}", getPayloadHandler)
		r.Delete("/payload/{id}", deletePayloadHandler)
		r.Post("/payload/{id}/view", trackViewHandler)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}