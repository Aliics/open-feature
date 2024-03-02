package database

import "errors"

type Database interface {
	All() ([]Flag, error)
	Get(key string) (*Flag, error)
	Put(flag Flag) error
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
	case "mem":
		return NewMemDatabase(), nil
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
	TypeMem  = "mem"
)

var (
	ErrUnknownDatabaseType = errors.New("unknown database type")
)
