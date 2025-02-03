package databasedriver

import (
	"context"
	"social/internal/models"
	"social/internal/store"
	socialv1 "social/proto/gen/social/v1"

	"github.com/gofrs/uuid"

	"gorm.io/gorm"
)

var _ store.Store = (*Store)(nil)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) NewTx(ctx context.Context) (store.Store, error) {
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &Store{db: tx}, nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) Commit() error {
	return s.db.Commit().Error
}

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
