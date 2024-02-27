package database

import "errors"

type Database interface {
	Get(key string) (bool, error)
	Put(key string, rule Rule) error
	Delete(key string) error
}

func NewDatabase(
	databaseType string,
	port int,
	host, dbname, user, password string,
) (Database, error) {
	switch databaseType {
	case "psql":
		return NewPSQLDatabase(port, host, dbname, user, password)
	default:
		return nil, ErrUnknownDatabaseType
	}
}

const (
	DefaultPort     = 5432
	DefaultHost     = "localhost"
	DefaultDBName   = "open-feature"
	DefaultUser     = "postgres"
	DefaultPassword = "password"

	TypePSQL = "psql"
)

var (
	ErrUnknownDatabaseType = errors.New("unknown database type")
)
