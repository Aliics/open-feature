package database

type Database interface {
	Get(key string) (bool, error)
	Put(key string, rule Rule) error
	Delete(key string) error
}

func NewDatabase(databaseType string, port int) (Database, error) {
	return nil, nil
}
