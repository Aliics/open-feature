package result

type Resulter interface {
	Result(v any) error
}
