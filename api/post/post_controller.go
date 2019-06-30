package post

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
)

type controller struct {
	repository           *storage.Repository
	credentialsExtractor web.CredentialsExtractor
}

func (c *controller) add(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not extract body: %s\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	username, _, err := c.credentialsExtractor.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repository.AddPost(ctx, username, body)
	if err != nil {
		log.Printf("Persisting post for user %s failed: %s\n", username, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (c *controller) get(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username, _, err := c.credentialsExtractor.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	posts, err := c.repository.GetTargetsPosts(ctx, username)
	if err != nil {
		log.Printf("Get posts for %s failed: %s\n", username, err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	web.WriteResponse(rw, http.StatusOK, posts)
}
