package user

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/cc2k19/go-tin/models"

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

	err = c.repository.AddUser(ctx, body)
	if err != nil {
		log.Printf("Persisting user failed: %s\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (c *controller) getByUsername(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := strings.TrimPrefix(r.URL.Path, web.UsersURL+"/")

	user, err := c.repository.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Get user with username %s failed: %s", username, err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	web.WriteResponse(rw, http.StatusOK, user)
}

func (c *controller) follow(rw http.ResponseWriter, r *http.Request) {
	err := c.executeRelation(r, c.repository.AddFollowRecord)
	if err != nil {
		log.Printf("Persisting unfollow relation failed: %s\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (c *controller) unfollow(rw http.ResponseWriter, r *http.Request) {
	err := c.executeRelation(r, c.repository.DeleteFollowRecord)
	if err != nil {
		log.Printf("Persisting unfollow relation failed: %s\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (c *controller) executeRelation(r *http.Request, f func(ctx context.Context, follower string, target string) error) error {
	defer r.Body.Close()

	ctx := r.Context()

	username, _, err := c.credentialsExtractor.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		return err
	}

	target := strings.TrimPrefix(r.URL.Path, web.FollowURL+"/")

	return f(ctx, username, target)
}

func (c *controller) getFollowers(rw http.ResponseWriter, r *http.Request) {
	c.getInteraction(rw, r, c.repository.GetFollowers)
}

func (c *controller) getFollowing(rw http.ResponseWriter, r *http.Request) {
	c.getInteraction(rw, r, c.repository.GetFollowing)
}

func (c *controller) getInteraction(rw http.ResponseWriter, r *http.Request, f func(ctx context.Context, username string) (models.UserSlice, error)) {
	ctx := r.Context()

	username, _, err := c.credentialsExtractor.Extract(r)
	if err != nil {
		log.Printf("Authorization decode error: %s", err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	followers, err := f(ctx, username)
	if err != nil {
		log.Printf("Get followers for user %s failed: %s", username, err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if followers == nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	web.WriteResponse(rw, http.StatusOK, followers)
}
