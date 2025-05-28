package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(ConfigureCORS())
	r.Use(NewRateLimiter(1, 3).EnforceRateLimit())

	r.Get("/", handleRoot)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/payload", createPayloadHandler)
		r.Get("/payload/{id}", getPayloadHandler)
		r.Delete("/payload/{id}", deletePayloadHandler)
		r.Get("/health", checkHealth)
})

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}