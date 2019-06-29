package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/cc2k19/go-tin/models"
)

type Repository struct {
	storage Storage
}

func NewRepository(storage Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (r *Repository) AddUser(ctx context.Context, user []byte) error {
	u := models.User{}
	if err := json.Unmarshal(user, &u); err != nil {
		return fmt.Errorf("could not parse json into user struct: %s", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not generate encrypted password")
	}

	u.Password = string(hashedPassword)

	return u.Insert(ctx, r.storage.Get(), boil.Infer())
}

func (r *Repository) AssertCredentials(ctx context.Context, username, password []byte) error {
	user, err := r.getUserByUsername(ctx, string(username))
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), password)
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) AddFollowRecord(ctx context.Context, follower string, target string) error {
	users, err := r.getFollowerTargetPair(ctx, follower, target)
	if err != nil {
		return err
	}

	followerUser, targetUser := users[0], users[1]

	return targetUser.AddFollowerUsers(ctx, r.storage.Get(), false, followerUser)
}

func (r *Repository) DeleteFollowRecord(ctx context.Context, follower string, target string) error {
	users, err := r.getFollowerTargetPair(ctx, follower, target)
	if err != nil {
		return err
	}

	followerUser, targetUser := users[0], users[1]

	return targetUser.RemoveTargetUsers(ctx, r.storage.Get(), followerUser)
}

func (r *Repository) GetFollowers(ctx context.Context, username, funcType string) (models.UserSlice, error) {
	users, err := r.getOneToMany(ctx, username, "Followers")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetFollowing(ctx context.Context, username string) (models.UserSlice, error) {
	users, err := r.getOneToMany(ctx, username, "Targets")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) AddPost(ctx context.Context, username string, post []byte) error {
	p := models.Post{}
	if err := json.Unmarshal(post, &p); err != nil {
		return fmt.Errorf("could not parse json into user struct")
	}

	p.Date = time.Now()

	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	p.UserID = user.ID

	return p.Insert(ctx, r.storage.Get(), boil.Infer())
}

func (r *Repository) GetTargetsPosts(ctx context.Context, username string) (models.PostSlice, error) {
	users, err := r.getOneToMany(ctx, username, "Targets")
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(users))
	for _, u := range users {
		ids = append(ids, u.ID)
	}

	return models.Posts(qm.WhereIn("user_id in ?", ids)).All(ctx, r.storage.Get())
}

func (r *Repository) getOneToMany(ctx context.Context, username, relation string) (models.UserSlice, error) {
	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var users models.UserSlice
	var e error

	switch relation {
	case "Followers":
		users, e = user.FollowerUsers(nil).All(ctx, r.storage.Get())
	case "Targets":
		users, e = user.TargetUsers(nil).All(ctx, r.storage.Get())
	}

	if e != nil {
		return nil, e
	}

	return users, nil
}

func (r *Repository) getUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := models.Users(qm.Where("username=?", username)).One(ctx, r.storage.Get())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) getFollowerTargetPair(ctx context.Context, follower string, target string) ([]*models.User, error) {
	targetUser, err := r.getUserByUsername(ctx, target)
	if err != nil {
		return nil, err
	}

	followerUser, err := r.getUserByUsername(ctx, follower)
	if err != nil {
		return nil, err
	}

	return []*models.User{followerUser, targetUser}, nil
}
