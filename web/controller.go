package web

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

// Controller is an entity that wraps a set of HTTP Routes
type Controller interface {
	// Routes returns the set of routes for this controller
	Routes() []Route
}

type AuthType string

const (
	BasicAuthentication AuthType = "basic"
	NoAuthentication    AuthType = "none"
)

// Route is a mapping between an Endpoint and a REST API Handler
type Route struct {
	// Endpoint is the combination of Path and HTTP Method for the specified route
	Endpoint Endpoint

	// Handler is the function that should handle incoming requests for this endpoint
	Handler http.HandlerFunc

	// Authentication mechanism for the route
	AuthType AuthType
}

// Endpoint is a combination of a Path and an HTTP Method
type Endpoint struct {
	Method string
	Path   string
}

type CredentialsExtractor interface {
	Extract(*http.Request) (string, error)
}

type FuncCredentialsExtractor func(*http.Request) (string, error)

func (fce FuncCredentialsExtractor) Extract(r *http.Request) (string, error) {
	return fce(r)
}

func BasicCredentialsExtract(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		return "", err
	}

	return strings.Split(string(decodedCredentials), ":")[0], nil
}
