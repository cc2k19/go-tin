package user

import (
	"github.com/cc2k19/go-tin/storage"
	"net/http"
)

type controller struct {
	repository *storage.Repository
}

func (c *controller) add(wr http.ResponseWriter, r *http.Request) {

}

func (c *controller) getByUsername(wr http.ResponseWriter, r *http.Request) {

}

func (c *controller) follow(wr http.ResponseWriter, r *http.Request) {

}

func (c *controller) unfollow(wr http.ResponseWriter, r *http.Request) {

}

func (c *controller) getFollowers(wr http.ResponseWriter, r *http.Request) {

}

func (c *controller) getFollowing(wr http.ResponseWriter, r *http.Request) {

}
