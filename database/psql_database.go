package database

type PSQLDatabase struct {
}

func NewPSQLDatabase(port int) (*PSQLDatabase, error) {
	//TODO implement me
	panic("implement me")
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
