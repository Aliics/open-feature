package result

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"reflect"
)

// Write will attempt to map the argument v to the Resulter given as a type parameter.
// If v is a slice or an array, it will do this operation iteratively.
// Once mapped, we will write this to the http.Response.
func Write[M Resulter](w http.ResponseWriter, v any) {
	one := func(v any) (M, error) {
		m := reflect.New(reflect.TypeFor[M]().Elem()).Interface().(M)
		err := m.Result(v)
		return m, err
	}

	var res any
	{
		vReflectValue := reflect.ValueOf(v)
		switch vReflectValue.Type().Kind() {
		case reflect.Slice, reflect.Array:
			mapped := make([]M, vReflectValue.Len())
			for i := range mapped {
				m, err := one(vReflectValue.Index(i).Interface())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				mapped[i] = m
			}
			res = mapped
		default:
			var err error
			res, err = one(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	WriteAny(w, res)
}

// WriteAny just wraps a generic json.Encoder call, but if a failure occurs it just emits a warning log.
// Instead of "handling this properly", we put our faith in Go's ability to write some JSON as a response.
// If for some reason it failed, we don't lose the error entirely.
func WriteAny(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		slog.Warn("writing json silently failed", "err", err)
	}
}
