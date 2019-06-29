package filters

import (
	"encoding/base64"
	"fmt"
	"github.com/cc2k19/go-tin/web"
	"log"
	"net/http"
	"strings"
)

type BasicAuthenticationFilter struct {
	repository *repository.Repository
}

func NewBasicAuthenticationFilter(repository *repository.Repository) *BasicAuthenticationFilter {
	return &BasicAuthenticationFilter{
		repository: repository,
	}
}

func (ba *BasicAuthenticationFilter) Filter(r *http.Request) (int, error) {
	authError := fmt.Errorf("authentication failed")

	ctx := r.Context()

	auth := r.Header.Get("Authorization")

	if !strings.HasPrefix(auth, "Basic ") {
		log.Println("Invalid authorization header")
		return http.StatusUnauthorized, authError
	}

	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		return http.StatusUnauthorized, authError
	}

	credentials := strings.Split(string(decodedCredentials), ":")

	err = ba.repository.AssertCredentials(ctx, []byte(credentials[0]), []byte(credentials[1]))
	if err != nil {
		log.Printf("credentials missmatch: %s", err)
		return http.StatusUnauthorized, authError
	}

	return http.StatusOK, nil
}

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
