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

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), mux)
}
