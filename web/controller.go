package web

import (
	"net/http"
)

// Controller is an entity that wraps a set of HTTP Routes
type Controller interface {
	// Routes returns the set of routes for this controller
	Routes() []Route
}

// HandlerFunc is an adapter that allows to use regular functions as Handler interface implementations.
type HandlerFunc func(req *http.Request) (resp *http.Response, err error)

// Route is a mapping between an Endpoint and a REST API Handler
type Route struct {
	// Endpoint is the combination of Path and HTTP Method for the specified route
	Endpoint Endpoint

	// Handler is the function that should handle incoming requests for this endpoint
	Handler HandlerFunc
}

// Endpoint is a combination of a Path and an HTTP Method
type Endpoint struct {
	Method string
	Path   string
}
