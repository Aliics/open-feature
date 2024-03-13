package api

import (
	"errors"
	"net/http"
	"open-feature/api/result"
	"open-feature/database"
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
	key := r.PathValue(PathValueKey)
	flag, err := s.Database.Get(key)
	if err != nil {
		if errors.Is(err, database.ErrFlagNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result.Write[*result.Flag](w, flag)
}

func (s *Server) createFlag(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) updateFlag(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) deleteFlag(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue(PathValueKey)

	err := s.Database.Delete(key)
	if err != nil {
		if errors.Is(err, database.ErrFlagNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const (
	PathValueKey = "key"
)
