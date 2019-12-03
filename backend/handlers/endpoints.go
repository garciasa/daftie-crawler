package handlers

import "github.com/go-chi/chi"

func (s *Server) setupEndPoints(r *chi.Mux) {
	r.Route("api/v1", func(r chi.Router) {
		r.Route("/houses", func(r chi.Router) {

		})
	})
}
