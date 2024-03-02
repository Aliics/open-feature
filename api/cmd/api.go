package main

import (
	"flag"
	"log/slog"
	"open-feature/api"
	"open-feature/database"
)

var (
	port             = flag.Int("port", api.DefaultPort, "specify the port the api listens on")
	databaseType     = flag.String("database-type", database.TypePSQL, "specify the backend database type to use")
	databasePort     = flag.Int("database-port", database.DefaultPort, "specify the backend database port to connect to")
	databaseHost     = flag.String("database-host", database.DefaultHost, "specify the backend database host")
	databaseDBName   = flag.String("database-dbname", database.DefaultDBName, "specify the backend database dbname")
	databaseUser     = flag.String("database-user", database.DefaultUser, "specify the backend database user")
	databasePassword = flag.String("database-password", database.DefaultPassword, "specify the backend database password")
)

func main() {
	flag.Parse()
	db, err := database.NewDatabase(
		*databaseType,
		*databasePort,
		*databaseDBName,
		*databaseHost,
		*databaseUser,
		*databasePassword,
	)
	if err != nil {
		slog.Error("database creation failed", "err", err)
	}

	s := api.Server{
		Config: api.Config{
			Port: *port,
		},
		Database: db,
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
