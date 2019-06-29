package api

import (
	"encoding/json"
	"github.com/cc2k19/go-tin/api/filters"
	"github.com/cc2k19/go-tin/api/post"
	"github.com/cc2k19/go-tin/api/user"
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
	"net/http"
)

func New(repository storage.Repository) *web.API {
	return &web.API{
		Controllers: []web.Controller{
			user.NewUsersController(repository),
			post.NewPostsController(repository),
		},
		Filters: []web.Filter{
			filters.NewBasicAuthenticationFilter(repository),
		},
	}
}

// ErrorResponse defines an error response payload
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteResponse writes a payload to the provided writer
func WriteResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
