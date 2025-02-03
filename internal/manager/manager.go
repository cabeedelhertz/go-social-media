package manager

import (
	"context"
	"fmt"
	socialv1 "social/proto/gen/social/v1"

	"social/internal/store"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type Manager struct {
	store      store.Store
	imageStore store.ImageStore
}

func NewManager(store store.Store, imageStore store.ImageStore) *Manager {
	return &Manager{store: store, imageStore: imageStore}
}

func (m *Manager) CreatePost(ctx context.Context, req *socialv1.CreatePostRequest) (*socialv1.CreatePostResponse, error) {
	post, err := m.store.CreatePost(ctx, req)
	if err != nil {
		return nil, err
	}
	return &socialv1.CreatePostResponse{Post: post}, nil
}

func (m *Manager) UpdatePost(ctx context.Context, req *socialv1.UpdatePostRequest) (*socialv1.UpdatePostResponse, error) {
	post, err := m.store.UpdatePostCaption(ctx, req.GetId(), req.GetCaption())
	if err != nil {
		return nil, err
	}
	return &socialv1.UpdatePostResponse{Post: post}, nil
}

func buildObjectKey(postId string) string {
	return fmt.Sprintf("posts/%s/%v.jpg", postId, uuid.Must(uuid.NewV4()))
}

func (m *Manager) AddImage(ctx context.Context, req *socialv1.AddImageRequest) (*socialv1.AddImageResponse, error) {
	if !isJPEGMagicBytes(req.Image.GetContent()) {
		return nil, errors.New("invalid image format")
	}
	objectKey := buildObjectKey(req.PostId)

	tx, err := m.store.NewTx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback()

	_, err = tx.AddImage(ctx, &socialv1.PostImage{
		PostId:    req.GetPostId(),
		ObjectKey: objectKey,
	})
	if err != nil {
		return nil, err
	}

	err = m.imageStore.UploadImage(ctx, objectKey, req.Image.GetContent())
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "failed to commit transaction")
	}

	post, err := m.store.GetPost(ctx, req.GetPostId())
	if err != nil {
		return nil, err
	}

	return &socialv1.AddImageResponse{Post: post}, nil

}

func (m *Manager) PublishPost(ctx context.Context, req *socialv1.PublishPostRequest) (*socialv1.PublishPostResponse, error) {
	post, err := m.store.PublishPost(ctx, req.GetPostId())
	if err != nil {
		return nil, err
	}
	return &socialv1.PublishPostResponse{Post: post}, nil
}

func (m *Manager) GetPost(ctx context.Context, req *socialv1.GetPostRequest) (*socialv1.GetPostResponse, error) {
	post, err := m.store.GetPost(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &socialv1.GetPostResponse{Post: post}, nil
}

func (m *Manager) GetPostImage(ctx context.Context, req *socialv1.GetPostImageRequest) (*socialv1.GetPostImageResponse, error) {
	postImage, err := m.store.GetPostImage(ctx, req.GetId(), req.GetPostId())
	if err != nil {
		return nil, err
	}

	imageUrl, err := m.imageStore.GetImageURL(ctx, postImage.GetObjectKey())
	if err != nil {
		return nil, err
	}

	postImage.Url = imageUrl
	return &socialv1.GetPostImageResponse{Image: postImage}, nil
}

func (m *Manager) ListPosts(ctx context.Context, req *socialv1.ListPostsRequest) (*socialv1.ListPostsResponse, error) {
	posts, err := m.store.ListPosts(ctx, req)
	if err != nil {
		return nil, err
	}
	return &socialv1.ListPostsResponse{Posts: posts}, nil
}

func (m *Manager) LikePost(ctx context.Context, req *socialv1.LikePostRequest) (*socialv1.LikePostResponse, error) {
	if req.UserId == "" || req.PostId == "" {
		return nil, errors.New("invalid input")
	}
	post, err := m.store.LikePost(ctx, req)
	if err != nil {
		return nil, err
	}
	return &socialv1.LikePostResponse{Post: post}, nil
}

func (m *Manager) HasUserLikedPost(ctx context.Context, postId string, userId string) (bool, error) {
	if postId == "" || userId == "" {
		return false, errors.New("invalid input")
	}
	return m.store.HasUserLikedPost(ctx, postId, userId)
}

func (m *Manager) UnlikePost(ctx context.Context, req *socialv1.UnlikePostRequest) (*socialv1.UnlikePostResponse, error) {
	if req.UserId == "" || req.PostId == "" {
		return nil, errors.New("invalid input")
	}
	post, err := m.store.UnlikePost(ctx, req)
	if err != nil {
		return nil, err
	}
	return &socialv1.UnlikePostResponse{Post: post}, nil
}

func isJPEGMagicBytes(data []byte) bool {
	// JPEG commonly starts with 0xFF 0xD8 and ends with 0xFF 0xD9,
	// but the most critical check is the start.
	if len(data) < 4 {
		return false
	}

	// Check the first two bytes
	if data[0] == 0xFF && data[1] == 0xD8 {
		// Optionally check the last two bytes, but
		// some images may have trailing data or custom segments
		// so it might be best to rely primarily on the first two.
		return true
	}
	return false
}
