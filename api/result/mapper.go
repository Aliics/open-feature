package result

type Mapper interface {
	Map(v any) error
}
