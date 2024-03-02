package database

import (
	"errors"
	"slices"
)

type MemDatabase struct {
	flags []Flag
}

func NewMemDatabase() *MemDatabase { return &MemDatabase{} }

func (m *MemDatabase) All() ([]Flag, error) {
	return m.flags, nil
}

func (m *MemDatabase) Get(key string) (*Flag, error) {
	for _, f := range m.flags {
		if f.Key != key {
			continue
		}
		return &f, nil
	}

	return nil, ErrFlagNotFound
}

func (m *MemDatabase) Put(flag Flag) error {
	i := slices.IndexFunc(m.flags, flagByKey(flag.Key))
	if i < 0 {
		m.flags = append(m.flags, flag)
		return nil
	}

	m.flags[i] = flag

	return nil
}

func (m *MemDatabase) Delete(key string) error {
	m.flags = slices.DeleteFunc(m.flags, flagByKey(key))
	return nil
}

func flagByKey(key string) func(flag Flag) bool {
	return func(flag Flag) bool { return flag.Key == key }
}

var (
	ErrFlagNotFound = errors.New("could not find flag with key")
)
