package web

const (
	apiVersion = "v1"

	// UsersURL is the URL path to manage Users
	UsersURL = "/" + apiVersion + "/users"

	// PostsURL is the URL path to manage Posts
	PostsURL = "/" + apiVersion + "/posts"

	// FollowURL is the URL path to perform following/unfollowing actions
	FollowURL = "/" + apiVersion + "/follow"

	// FollowersURL is the URL path to manage your followers
	FollowersURL = "/" + apiVersion + "/followers"

	// FollowingURL is the URL path to manage users you follow
	FollowingURL = "/" + apiVersion + "/following"
)
