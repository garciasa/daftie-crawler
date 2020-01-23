package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

func (s *Server) setupEndPoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/houses", func(r chi.Router) {
			r.Get("/", s.getAllHouses())
		})
	})
	r.Route("/app", func(r chi.Router) {
		workDir, _ := os.Getwd()
		log.Printf("directory: %s", workDir)
		staticDir := filepath.Join(workDir, "build")
		FileServer(r, "/app", "/static", http.Dir(staticDir))
	})
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
