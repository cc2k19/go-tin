package user

import (
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
	"net/http"
)

func NewUsersController(repository *storage.Repository) *controller {
	return &controller{
		repository: repository,
	}
}

func (c *controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   web.UsersURL,
			},
			AuthType: web.NoAuthentication,
			Handler:  c.add,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.UsersURL + "/{username}",
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.getByUsername,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPut,
				Path:   web.FollowURL + "/{username}",
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.follow,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   web.FollowURL + "/{username}",
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.unfollow,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.FollowersURL,
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.getFollowers,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.FollowingURL,
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.getFollowing,
		},
	}
}
