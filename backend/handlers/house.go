package handlers

import "net/http"

func (s *Server) getAllHouses() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		houses, err := s.domain.GetAll()
		if err != nil{
			badRequestResponse(w, err)
		}

		jsonResponse(w, houses, http.StatusOK)
	}
}