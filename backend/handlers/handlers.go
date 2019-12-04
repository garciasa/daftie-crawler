package handlers

import "github.com/go-chi/chi"
import "github.com/go-chi/chi/middleware"
import "time"
// Server api server struct
type Server struct {
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))

}
// NewServer create a new api server instance
func NewServer() *Server {
	return &Server{}
}

// SetupRouter set up all routes
func SetupRouter() *chi.Mux {
	server := NewServer()

	r := chi.NewRouter()

	server.setupEndPoints(r)

	return r
}
