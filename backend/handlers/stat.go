package handlers

import (
	"errors"
	"net/http"
	"strconv"
)

func (s *Server) getStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := s.domain.GetStats()
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, stats, http.StatusOK)
	}
}

func (s *Server) getStatsForCharts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("beds")
		
		if param == "" {

			badRequestResponse(w, errors.New("Error with parameters"))
			return
		}

		beds, err := strconv.Atoi(param)
		if err != nil {
			badRequestResponse(w, err)
			return
		}
		
		provider := r.URL.Query().Get("provider")
		stats, err := s.domain.GetStatsForCharts(provider, beds)
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, stats, http.StatusOK)
	}
}