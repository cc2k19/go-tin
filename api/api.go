package api

import (
	"github.com/cc2k19/go-tin/api/filters"
	"github.com/cc2k19/go-tin/api/post"
	"github.com/cc2k19/go-tin/api/user"
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

// New returns new web API with controllers and filters for given repository
func New(repository *storage.Repository) *web.API {
	return &web.API{
		Controllers: []web.Controller{
			user.NewUsersController(repository, web.CredentialsExtractorFunc(web.BasicCredentialsExtractor)),
			post.NewPostsController(repository, web.CredentialsExtractorFunc(web.BasicCredentialsExtractor)),
		},
		Filters: []web.Filter{
			filters.NewBasicAuthenticationFilter(repository),
		},
	}
}
