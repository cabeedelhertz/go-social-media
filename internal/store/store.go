package store

import (
	"context"
	socialv1 "social/proto/gen/social/v1"

	"github.com/gofrs/uuid"
)

type Store interface {
	NewTx(ctx context.Context) (Store, error)
	Rollback() error
	Commit() error
	CreatePost(ctx context.Context, req *socialv1.CreatePostRequest) (*socialv1.Post, error)
	UpdatePostCaption(ctx context.Context, id string, caption string) (*socialv1.Post, error)
	AddImage(ctx context.Context, pi *socialv1.PostImage) (*socialv1.PostImage, error)
	PublishPost(ctx context.Context, id string) (*socialv1.Post, error)
	GetPost(ctx context.Context, id string) (*socialv1.Post, error)
	GetPostImage(ctx context.Context, id string, postId string) (*socialv1.PostImage, error)
	ListPosts(ctx context.Context, req *socialv1.ListPostsRequest) ([]*socialv1.Post, error)
	CreateOrUpdateUser(ctx context.Context, usr *socialv1.User) (*socialv1.User, error)
	GetUser(ctx context.Context, id string) (*socialv1.User, error)
	LikePost(ctx context.Context, req *socialv1.LikePostRequest) (*socialv1.Post, error)
	UnlikePost(ctx context.Context, req *socialv1.UnlikePostRequest) (*socialv1.Post, error)
	HasUserLikedPost(ctx context.Context, postId string, userId string) (bool, error)
	ListUsers(ctx context.Context, req *socialv1.ListUsersRequest) ([]*socialv1.User, error)
	CreateFollower(ctx context.Context, userId, followerId uuid.UUID) (*socialv1.Follower, error)
	DeleteFollower(ctx context.Context, userId, followerId uuid.UUID) error
	ListFollowers(ctx context.Context, req *socialv1.ListFollowersRequest) ([]*socialv1.Follower, error)
	GetIsFollowing(ctx context.Context, userId, followerId uuid.UUID) (bool, error)
}

type ImageStore interface {
	UploadImage(ctx context.Context, key string, content []byte) error
	GetImageURL(ctx context.Context, key string) (string, error)
}
