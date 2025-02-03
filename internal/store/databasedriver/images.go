package databasedriver

import (
	"context"
	"social/internal/models"
	socialv1 "social/proto/gen/social/v1"

	"github.com/gofrs/uuid"
)

func (s *Store) AddImage(ctx context.Context, pi *socialv1.PostImage) (*socialv1.PostImage, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).Preload("Images").Where("id = ?", pi.PostId).First(&post).Error; err != nil {
		return nil, err
	}
	image := models.PostImage{
		ID:         uuid.Must(uuid.NewV4()),
		PostID:     post.ID,
		ObjectKey:  pi.GetObjectKey(),
		OrderIndex: len(post.Images),
	}
	if err := s.db.WithContext(ctx).Create(&image).Error; err != nil {
		return nil, err
	}
	return image.ToProto(), nil
}

func (s *Store) GetPostImage(ctx context.Context, id string, postId string) (*socialv1.PostImage, error) {
	var image models.PostImage
	if err := s.db.WithContext(ctx).Where("id = ? AND post_id = ?", id, postId).First(&image).Error; err != nil {
		return nil, err
	}
	return image.ToProto(), nil
}
