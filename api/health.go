package api

import (
	"log/slog"
	"net/http"
	"open-feature/database"
)

type healthCheckResponse struct {
	DB status `json:"db"`
}

type status string

const (
	statusOK  status = "OK"
	statusERR status = "ERR"
)

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	var res healthCheckResponse

	{
		switch db := s.Database.(type) {
		case *database.PSQLDatabase:
			if err := db.DB.Ping(); err != nil {
				slog.Error("psql connectivity errored on healthcheck", "err", err)
				res.DB = statusERR
			} else {
				res.DB = statusOK
			}
		default:
			// No connectivity needed.
			res.DB = statusOK
		}
	}

	writeJSON(w, res)
}
