package user

import (
	"github.com/cc2k19/go-tin/api"
	"github.com/cc2k19/go-tin/storage"
	"io/ioutil"
	"log"
	"net/http"
)

type controller struct {
	repository *storage.Repository
}

func (c *controller) add(wr http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bodyReader, err := r.GetBody()
	if err != nil {
		log.Printf("Could not extract body: %s\n", err)
		api.WriteResponse(wr, http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
	}
	defer bodyReader.Close()

	body, err := ioutil.ReadAll(bodyReader)

	if err != nil {
		log.Printf("Could not extract body: %s\n", err)
		api.WriteResponse(wr, http.StatusBadRequest, api.ErrorResponse{Error: err.Error()})
	}

	users, err := c.repository.AddUser(ctx, body)
	if err != nil {
		log.Printf("Persisting user failed: %s\n", err)
	}
	api.WriteResponse(wr, http.StatusBadRequest, users)
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
