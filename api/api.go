package api

import (
	"github.com/cc2k19/go-tin/api/post"
	"github.com/cc2k19/go-tin/api/user"
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

func New(storage storage.Storage) *web.API {
	return &web.API{Controllers: []web.Controller{
		user.NewUsersController(),
		post.NewPostsController(),
	}}
}
