package input

import (
	"encoding/json"
	"net/http"
)

func GetValidatedInput[V Validator](r *http.Request) (*V, error) {
	var v V

	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return nil, err
	}

	err = v.Validate()
	if err != nil {
		return nil, err
	}

	return &v, nil
}
