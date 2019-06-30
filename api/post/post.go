package post

import (
	"net/http"

	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

func NewPostsController(repository *storage.Repository, extractor web.CredentialsExtractor) *controller {
	return &controller{
		repository:           repository,
		credentialsExtractor: extractor,
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
