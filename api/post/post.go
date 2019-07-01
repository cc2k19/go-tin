package post

import (
	"net/http"

	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

// NewPostsController returns new post controller for given repository and mechanism for credentials extracting from request
func NewPostsController(repository storage.Repository, extractor web.CredentialsExtractor) *controller {
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
