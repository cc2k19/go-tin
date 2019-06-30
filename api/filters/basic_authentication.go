package filters

import (
	"fmt"
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
	"log"
	"net/http"
)

// BasicAuthenticationFilter provides security with basic authentication mechanism
type BasicAuthenticationFilter struct {
	repository           *storage.Repository
	credentialsExtractor web.CredentialsExtractor
}

// NewBasicAuthenticationFilter returns new basic auth filter for given repository
func NewBasicAuthenticationFilter(repository *storage.Repository) *BasicAuthenticationFilter {
	return &BasicAuthenticationFilter{
		repository:           repository,
		credentialsExtractor: web.CredentialsExtractorFunc(web.BasicCredentialsExtractor),
	}
}

// Filter filters http request on some conditions
func (ba *BasicAuthenticationFilter) Filter(r *http.Request) (int, error) {
	authError := fmt.Errorf("authentication failed")

	ctx := r.Context()

	username, password, err := ba.credentialsExtractor.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		return http.StatusUnauthorized, authError
	}

	err = ba.repository.AssertCredentials(ctx, []byte(username), []byte(password))
	if err != nil {
		log.Printf("credentials missmatch: %s", err)
		return http.StatusUnauthorized, authError
	}

	return http.StatusOK, nil
}

// MatchingEndpoints returns all the endpoints that the filter should be attached before.
func (ba *BasicAuthenticationFilter) MatchingEndpoints() []web.Endpoint {
	return []web.Endpoint{
		{
			Method: http.MethodGet,
			Path:   web.UsersURL + "/{username}",
		},
		{
			Method: http.MethodPost,
			Path:   web.FollowURL + "/{username}",
		},
		{
			Method: http.MethodDelete,
			Path:   web.FollowURL + "/{username}",
		},
		{
			Method: http.MethodGet,
			Path:   web.FollowersURL,
		},
		{
			Method: http.MethodGet,
			Path:   web.FollowingURL,
		},
		{
			Method: http.MethodPost,
			Path:   web.PostsURL,
		},
		{
			Method: http.MethodGet,
			Path:   web.PostsURL,
		},
	}
}
