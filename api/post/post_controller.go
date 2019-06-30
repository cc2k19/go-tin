package post

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

type controller struct {
	repository *storage.Repository
	ce         web.CredentialsExtractor
}

func (c *controller) add(wr http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not extract body: %s\n", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	username, err := c.ce.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repository.AddPost(ctx, username, body)
	if err != nil {
		log.Printf("Persisting post for user %s failed: %s\n", username, err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}
	wr.WriteHeader(http.StatusCreated)
}

func (c *controller) get(wr http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username, err := c.ce.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	posts, err := c.repository.GetTargetsPosts(ctx, username)
	if err != nil {
		log.Printf("Get posts for %s failed: %s\n", username, err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}
	web.WriteResponse(wr, http.StatusOK, posts)
}
