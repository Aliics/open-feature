package api

import (
	"fmt"
	"net/http"
	"open-feature/database"
)

type Server struct {
	Config
	database.Database
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.Port),
		s.NewServeMux(),
	)
}

func (s *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /flags/all", s.listFlags)
	mux.HandleFunc("GET /flags/{key}", s.getFlag)

	return mux
}

const (
	DefaultPort = 8080
)
