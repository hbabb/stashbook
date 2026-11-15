package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// writeJSON sends data as JSON with the given status code.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("encode response", "err", err)
	}
}

// writeError sends {"error": message} as JSON with the given status code.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
