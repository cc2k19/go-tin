package web

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// Controller is an entity that wraps a set of HTTP Routes
type Controller interface {
	// Routes returns the set of routes for this controller
	Routes() []Route
}

// Route is a mapping between an Endpoint and a REST API Handler
type Route struct {
	// Endpoint is the combination of Path and HTTP Method for the specified route
	Endpoint Endpoint

	// Handler is the function that should handle incoming requests for this endpoint
	Handler http.HandlerFunc
}

// Endpoint is a combination of a Path and an HTTP Method
type Endpoint struct {
	Method string
	Path   string
}

// CredentialsExtractor is a generic interface for extracting credentials from different auth mechanism
type CredentialsExtractor interface {
	Extract(*http.Request) (string, string, error)
}

// CredentialsExtractorFunc is a wrapper of CredentialsExtractor so a normal function could be used as CredentialsExtractor
type CredentialsExtractorFunc func(*http.Request) (string, string, error)

// Extract extract the credentials and satisfy the CredentialsExtractor interface
func (cef CredentialsExtractorFunc) Extract(r *http.Request) (string, string, error) {
	return cef(r)
}

// BasicCredentialsExtractor extract credentials from basic auth header
func BasicCredentialsExtractor(r *http.Request) (string, string, error) {
	auth := r.Header.Get("Authorization")

	if !strings.HasPrefix(auth, "Basic ") {
		return "", "", fmt.Errorf("invalid authorization header")
	}

	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		return "", "", err
	}

	credentials := strings.Split(string(decodedCredentials), ":")

	return credentials[0], credentials[1], nil
}
