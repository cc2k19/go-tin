package post

import (
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
	"net/http"
)

func NewPostsController(repository *storage.Repository) *controller {
	return &controller{
		repository: repository,
	}
}

func (c *controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   web.PostsURL,
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.add,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.PostsURL,
			},
			AuthType: web.BasicAuthentication,
			Handler:  c.get,
		},
	}
}
