package api

import (
	"net/http"
	"open-feature/api/result"
)

func (s *Server) listFlags(w http.ResponseWriter, _ *http.Request) {
	flags, err := s.Database.All()
	if err != nil {
		// This has no user input effectively.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result.Write[*result.Flag](w, flags)
}

func (s *Server) getFlag(w http.ResponseWriter, r *http.Request) {

}
