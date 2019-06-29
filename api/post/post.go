package post

import (
	"github.com/cc2k19/go-tin/web"
	"net/http"
)

func NewPostsController() *controller {
	return &controller{}
}

func (c *controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   web.PostsURL,
			},
			Handler: c.add,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.PostsURL,
			},
			Handler: c.get,
		},
	}
}