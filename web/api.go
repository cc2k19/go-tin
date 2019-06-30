package web

import (
	"encoding/json"
	"net/http"
)

// API is the primary point for REST API registration
type API struct {
	// Controllers contains the registered controllers
	Controllers []Controller

	// Filters contains registered filters for this API
	Filters []Filter
}

// ErrorResponse defines an error response payload
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteResponse writes a payload to the provided writer
func WriteResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
