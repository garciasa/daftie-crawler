package handlers

import "net/http"

func (s *Server) getStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := s.domain.GetStats()
		if err != nil {
			badRequestResponse(w, err)
		}

		jsonResponse(w, stats, http.StatusOK)
	}
}
