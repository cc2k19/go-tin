package api

import (
	"github.com/cc2k19/go-tin/api/filters"
	"github.com/cc2k19/go-tin/api/post"
	"github.com/cc2k19/go-tin/api/user"
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

func New(repository *storage.Repository) *web.API {
	return &web.API{
		Controllers: []web.Controller{
			user.NewUsersController(repository, web.FuncCredentialsExtractor(web.BasicCredentialsExtract)),
			post.NewPostsController(repository, web.FuncCredentialsExtractor(web.BasicCredentialsExtract)),
		},
		Filters: []web.Filter{
			filters.NewBasicAuthenticationFilter(repository),
		},
	}
}
