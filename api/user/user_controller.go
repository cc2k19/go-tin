package user

import (
	"encoding/base64"
	"github.com/cc2k19/go-tin/storage"
	"github.com/cc2k19/go-tin/web"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repository.AddUser(ctx, body)
	if err != nil {
		log.Printf("Persisting user failed: %s\n", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}
	wr.WriteHeader(http.StatusCreated)
}

func (c *controller) getByUsername(wr http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := strings.TrimPrefix(r.URL.Path, web.UsersURL+"/")

	user, err := c.repository.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Get user with username %s failed: %s", username, err)
		wr.WriteHeader(http.StatusNotFound)
		return
	}

	web.WriteResponse(wr, http.StatusOK, user)
}

func (c *controller) follow(wr http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	auth := r.Header.Get("Authorization")
	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	username := strings.Split(string(decodedCredentials), ":")[0]
	target := strings.TrimPrefix(r.URL.Path, web.FollowURL+"/")

	err = c.repository.AddFollowRecord(ctx, username, target)
	if err != nil {
		log.Printf("Persisting follow relation failed: %s\n", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusCreated)
}

func (c *controller) unfollow(wr http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	auth := r.Header.Get("Authorization")
	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	username := strings.Split(string(decodedCredentials), ":")[0]
	target := strings.TrimPrefix(r.URL.Path, web.FollowURL+"/")

	err = c.repository.DeleteFollowRecord(ctx, username, target)
	if err != nil {
		log.Printf("Persisting follow relation failed: %s\n", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	wr.WriteHeader(http.StatusOK)
}

func (c *controller) getFollowers(wr http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	auth := r.Header.Get("Authorization")
	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	username := strings.Split(string(decodedCredentials), ":")[0]
	followers, err := c.repository.GetFollowers(ctx, username)
	if err != nil {
		log.Printf("Get followers for user %s failed: %s", username, err)
		wr.WriteHeader(http.StatusNotFound)
	}

	web.WriteResponse(wr, http.StatusOK, followers)
}

func (c *controller) getFollowing(wr http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	auth := r.Header.Get("Authorization")
	decodedCredentials, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	username := strings.Split(string(decodedCredentials), ":")[0]
	followers, err := c.repository.GetFollowing(ctx, username)
	if err != nil {
		log.Printf("Get followers for user %s failed: %s", username, err)
		wr.WriteHeader(http.StatusNotFound)
	}

	web.WriteResponse(wr, http.StatusOK, followers)
}
