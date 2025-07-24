package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// DecodeJSON reads JSON from request body and populates the result struct.
// Returns error if decoding fails.
func DecodeJSON(r *http.Request, result interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Prevent extra fields from silently being ignored

	if err := decoder.Decode(result); err != nil {
		return fmt.Errorf("failed to decode JSON body: %w", err)
	}

	// Check if there's extra data in body
	if decoder.More() {
		return errors.New("multiple JSON objects in body are not allowed")
	}
	return nil
}

// RespondJSON encodes result to JSON and writes to response
func RespondJSON(w http.ResponseWriter, code int, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if result == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}
