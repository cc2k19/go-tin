package web

import "net/http"

// Filter interface provides an interface for filtering http requests on some conditions
type Filter interface {

	// Filter filters http request on some conditions
	Filter(r *http.Request) (int, error)

	// MatchingEndpoints returns all the endpoints that the filter should be attached before.
	MatchingEndpoints() []Endpoint
}
