package api

import "net/http"

type flag struct {
	Key string `json:"key"`
}

func (s *Server) listFlags(w http.ResponseWriter, _ *http.Request) {
	flags, err := s.Database.All()
	if err != nil {
		// This has no user input effectively.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := make([]flag, len(flags))
	for i, f := range flags {
		result[i] = flag{f.Key}
	}

	writeJSON(w, result)
}

func (s *Server) getFlag(w http.ResponseWriter, r *http.Request) {

}
