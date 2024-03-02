package database

type Rule interface {
	Eval() bool
}

type StaticRule bool

func (s StaticRule) Eval() bool { return bool(s) }
