package user

import (
	"github.com/cc2k19/go-tin/web"
	"net/http"
)

func NewUsersController() *controller {
	return &controller{}
}

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
				Method: http.MethodPost,
				Path: web.FollowURL + "/{username}",
			},
			Handler: c.follow,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path: web.FollowURL + "/{username}",
			},
			Handler: c.unfollow,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path: web.FollowersURL,
			},
			Handler: c.getFollowers,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path: web.FollowingURL,
			},
			Handler: c.getFollowing,
		},
	}
}
