package api

import (
	"fmt"
	"github.com/aliics/open-feature/database"
	"net/http"
)

type Server struct {
	Config
	database.Database
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.health)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), mux)
}

const (
	DefaultPort = 8080
)
