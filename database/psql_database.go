package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PSQLDatabase struct {
	DB *sql.DB
}

func NewPSQLDatabase(port int, host, dbname, user, password string) (*PSQLDatabase, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"port=%d host=%s dbname=%s user=%s password=%s",
			port,
			host,
			dbname,
			user,
			password,
		),
	)
	if err != nil {
		return nil, err
	}

	return &PSQLDatabase{db}, nil
}

func (p *PSQLDatabase) Get(key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PSQLDatabase) Put(key string, rule Rule) error {
	//TODO implement me
	panic("implement me")
}

func (p *PSQLDatabase) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}
