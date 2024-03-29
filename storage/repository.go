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

// Repository wraps the storage and provides domain specific functions to interact with the storage
//go:generate counterfeiter . Repository
type Repository interface {
	AssertCredentials(ctx context.Context, username, password []byte) error
	AddUser(ctx context.Context, user []byte) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	AddFollowRecord(ctx context.Context, follower string, target string) error
	DeleteFollowRecord(ctx context.Context, follower string, target string) error
	GetFollowers(ctx context.Context, username string) (models.UserSlice, error)
	GetFollowing(ctx context.Context, username string) (models.UserSlice, error)
	AddPost(ctx context.Context, username string, post []byte) error
	GetTargetsPosts(ctx context.Context, username string) (models.PostSlice, error)
}

// RepositoryImpl implements Repository in order to interact with storage
type RepositoryImpl struct {
	storage Storage
}

// NewRepository returns new repository for a given storage
func NewRepository(storage Storage) Repository {
	return &RepositoryImpl{
		storage: storage,
	}
}

// AssertCredentials checks if given username and password are the same as persisted into the database
func (r *RepositoryImpl) AssertCredentials(ctx context.Context, username, password []byte) error {
	user, err := r.getUserByUsername(ctx, string(username))
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), password)
}

// AddUser adds user given as byte slice into the database
func (r *RepositoryImpl) AddUser(ctx context.Context, user []byte) error {
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

// GetUserByUsername returns user entity for a given user identified by username
func (r *RepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AddFollowRecord adds a follow relation between two users of type `follower follows target`
func (r *RepositoryImpl) AddFollowRecord(ctx context.Context, follower string, target string) error {
	users, err := r.getFollowerTargetPair(ctx, follower, target)
	if err != nil {
		return err
	}

	followerUser, targetUser := users[0], users[1]

	return targetUser.AddFollowerUsers(ctx, r.storage.Get(), false, followerUser)
}

// DeleteFollowRecord deletes a follow relation between two users
func (r *RepositoryImpl) DeleteFollowRecord(ctx context.Context, follower string, target string) error {
	users, err := r.getFollowerTargetPair(ctx, follower, target)
	if err != nil {
		return err
	}

	followerUser, targetUser := users[0], users[1]

	return followerUser.RemoveTargetUsers(ctx, r.storage.Get(), targetUser)
}

// GetFollowers returns all the followers of a given user
func (r *RepositoryImpl) GetFollowers(ctx context.Context, username string) (models.UserSlice, error) {
	users, err := r.getOneToMany(ctx, username, "Followers")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetFollowing returns all the users who follow a given user
func (r *RepositoryImpl) GetFollowing(ctx context.Context, username string) (models.UserSlice, error) {
	users, err := r.getOneToMany(ctx, username, "Targets")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// AddPost adds a post to a given users profile
func (r *RepositoryImpl) AddPost(ctx context.Context, username string, post []byte) error {
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

// GetTargetsPosts returns all the posts which a given user can see
func (r *RepositoryImpl) GetTargetsPosts(ctx context.Context, username string) (models.PostSlice, error) {
	users, err := r.getOneToMany(ctx, username, "Targets")
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, fmt.Errorf("current user does not follow anyone")
	}

	ids := make([]interface{}, 0, len(users))
	for _, u := range users {
		ids = append(ids, u.ID)
	}

	return models.Posts(qm.WhereIn("user_id in ?", ids...)).All(ctx, r.storage.Get())
}

func (r *RepositoryImpl) getOneToMany(ctx context.Context, username, relation string) (models.UserSlice, error) {
	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var emptyQueryMods []qm.QueryMod
	var users models.UserSlice
	var e error

	switch relation {
	case "Targets":
		users, e = user.TargetUsers(emptyQueryMods...).All(ctx, r.storage.Get())
	case "Followers":
		users, e = user.FollowerUsers(emptyQueryMods...).All(ctx, r.storage.Get())
	}

	if e != nil {
		return nil, e
	}

	return users, nil
}

func (r *RepositoryImpl) getUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := models.Users(qm.Where("username=?", username)).One(ctx, r.storage.Get())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *RepositoryImpl) getFollowerTargetPair(ctx context.Context, follower string, target string) ([]*models.User, error) {
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
