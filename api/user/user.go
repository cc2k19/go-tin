package user

import (
	"net/http"

	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

// NewUsersController returns new user controller for given repository and mechanism for credentials extracting from request
func NewUsersController(repository storage.Repository, extractor web.CredentialsExtractor) *controller {
	return &controller{
		repository:           repository,
		credentialsExtractor: extractor,
	}
}

// Routes returns the set of routes for this controller
func (c *controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   web.UsersURL,
			},
			Handler: c.add,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.UsersURL + "/{username}",
			},
			Handler: c.getByUsername,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPut,
				Path:   web.FollowURL + "/{username}",
			},
			Handler: c.follow,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   web.FollowURL + "/{username}",
			},
			Handler: c.unfollow,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.FollowersURL,
			},
			Handler: c.getFollowers,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.FollowingURL,
			},
			Handler: c.getFollowing,
		},
	}
}
