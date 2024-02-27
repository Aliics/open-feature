package database

import "errors"

type Database interface {
	Get(key string) (bool, error)
	Put(key string, rule Rule) error
	Delete(key string) error
}

func NewDatabase(databaseType string, port int) (Database, error) {
	switch databaseType {
	case "psql":
		return NewPSQLDatabase(port)
	default:
		return nil, ErrUnknownDatabaseType
	}
}

const (
	DefaultPort = 5432

	TypePSQL = "psql"
)

var (
	ErrUnknownDatabaseType = errors.New("unknown database type")
)
