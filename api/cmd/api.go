package main

import (
	"flag"
	"github.com/aliics/open-feature/api"
	"github.com/aliics/open-feature/database"
	"log/slog"
)

var (
	port         = flag.Int("port", api.DefaultPort, "specify the port the api listens on")
	databaseType = flag.String("database-type", database.TypePSQL, "specify the backend database type to use")
	databasePort = flag.Int("database-port", database.DefaultPort, "specify the backend database port to connect to")
)

func main() {
	flag.Parse()
	db, err := database.NewDatabase(*databaseType, *databasePort)
	if err != nil {
		slog.Error("database creation failed", "err", err)
	}

	s := api.Server{
		Config: api.Config{
			Port:     *port,
			Database: db,
		},
	}

	slog.Info("starting server", "port", *port)

	if err = s.ListenAndServe(); err != nil {
		slog.Error(
			"server failure",
			"port", s.Port,
			"err", err,
		)
	}
}
