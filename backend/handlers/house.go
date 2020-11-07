package handlers

import (
	"backend/domain"
	"net/http"
	"strconv"
)

func (s *Server) getAllHouses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var houses []domain.House
		var err error

		// provider := r.URL.Query().Get("provider")
		pag, err := strconv.Atoi(r.URL.Query().Get("pag"))
		
		if err != nil {
			houses, err = s.domain.GetAllHouses()
		} else {
			// houses, err = s.domain.GetHousesByProvider(provider)
			houses, err = s.domain.GetHousesPerPage(pag)
		}
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, houses, http.StatusOK)
	}
}

func (s *Server) getLastHouses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		houses, err := s.domain.GetLastHouses()
		if err != nil {
			badRequestResponse(w, err)
			return
		}

		jsonResponse(w, houses, http.StatusOK)
	}
}
