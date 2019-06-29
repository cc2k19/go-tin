package web

import "net/http"

type Filter interface {

	// Filter filters http request on some conditions
	Filter(r *http.Request) (int, error)

	MatchingEndpoints() []Endpoint
}
