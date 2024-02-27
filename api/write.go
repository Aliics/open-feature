package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// writeJSON just wraps a generic json.Encoder call, but if a failure occurs it just emits a warning log.
// Instead of "handling this properly", we put our faith in Go's ability to write some JSON as a response.
// If for some reason it failed, we don't lose the error entirely.
func writeJSON(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		slog.Warn("writing json silently failed", "err", err)
	}
}
