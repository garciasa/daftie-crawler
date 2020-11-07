package handlers

import (
	"net/http"
	"strconv"
)

func (s *Server) getStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := s.domain.GetStats()
		if err != nil {
			badRequestResponse(w, err)
		}

		jsonResponse(w, stats, http.StatusOK)
	}
}

func (s *Server) getStatsForCharts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		beds, err := strconv.Atoi(r.URL.Query().Get("beds"))
		if err != nil {
			badRequestResponse(w, err)
		}
		
		provider := r.URL.Query().Get("provider")
		stats, err := s.domain.GetStatsForCharts(provider, beds)
		if err != nil {
			badRequestResponse(w, err)
		}

		jsonResponse(w, stats, http.StatusOK)
	}
}