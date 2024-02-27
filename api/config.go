package api

import "github.com/aliics/open-feature/database"

type Config struct {
	Port     int
	Database database.Database
}
