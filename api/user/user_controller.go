package user

import (
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
	"io/ioutil"
	"log"
	"net/http"
)

type controller struct {
	repository *storage.Repository
}

func (c *controller) add(wr http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Could not extract body: %s\n", err)
		web.WriteResponse(wr, http.StatusBadRequest, web.ErrorResponse{Error: err.Error()})
		return
	}

	err = c.repository.AddUser(ctx, body)
	if err != nil {
		log.Printf("Persisting user failed: %s\n", err)
		web.WriteResponse(wr, http.StatusBadRequest, web.ErrorResponse{Error: err.Error()})
		return
	}
	wr.WriteHeader(http.StatusCreated)
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
