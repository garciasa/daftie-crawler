package handlers

import (
	"backend/domain"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// Server api server struct
type Server struct {
	domain *domain.Domain
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

}

// NewServer create a new api server instance
func NewServer(domain *domain.Domain) *Server {
	return &Server{domain: domain}
}

// SetupRouter set up all routes
func SetupRouter(domain *domain.Domain) *chi.Mux {
	server := NewServer(domain)

	r := chi.NewRouter()

	setupMiddleware(r)

	server.setupEndPoints(r)

	return r
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if data == nil {
		data = map[string]string{}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func badRequestResponse(w http.ResponseWriter, err error) {
	response := map[string]string{"error": err.Error()}
	jsonResponse(w, response, http.StatusBadRequest)
}
