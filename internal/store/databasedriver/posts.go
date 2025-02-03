package databasedriver

import (
	"context"
	"social/internal/models"
	socialv1 "social/proto/gen/social/v1"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (s *Store) CreatePost(ctx context.Context, req *socialv1.CreatePostRequest) (*socialv1.Post, error) {
	post := models.Post{
		ID:          uuid.Must(uuid.NewV4()),
		Description: req.Caption,
		UserID:      uuid.FromStringOrNil(req.UserId),
		Status:      models.PostStatusDraft,
	}
	if err := s.db.WithContext(ctx).Create(&post).Error; err != nil {
		return nil, err
	}
	return post.ToProto(), nil
}

func (s *Store) UpdatePostCaption(ctx context.Context, id string, caption string) (*socialv1.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	post.Description = caption
	if err := s.db.WithContext(ctx).Save(&post).Error; err != nil {
		return nil, err
	}
	return post.ToProto(), nil
}

func (s *Store) PublishPost(ctx context.Context, id string) (*socialv1.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	post.Status = models.PostStatusPublished
	if err := s.db.WithContext(ctx).Save(&post).Error; err != nil {
		return nil, err
	}
	return post.ToProto(), nil
}

func (s *Store) GetPost(ctx context.Context, id string) (*socialv1.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).Preload("Images").Preload("User").
		Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return post.ToProto(), nil
}

func (s *Store) ListPosts(ctx context.Context, req *socialv1.ListPostsRequest) ([]*socialv1.Post, error) {
	var posts []models.Post
	query := s.db.WithContext(ctx).Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("post_images.order_index ASC")
	}).Preload("User").Order("created_at DESC").
		Where("status = 'PUBLISHED'")
	if req.GetUserId() != "" {
		query = query.Where("user_id = ?", req.GetUserId())
	}
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	var resp []*socialv1.Post
	for _, post := range posts {
		resp = append(resp, post.ToProto())
	}
	return resp, nil
}

func (s *Store) LikePost(ctx context.Context, req *socialv1.LikePostRequest) (*socialv1.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).Where("id = ?", req.PostId).First(&post).Error; err != nil {
		return nil, err
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		like := models.PostLike{
			ID:     uuid.Must(uuid.NewV4()),
			PostID: uuid.FromStringOrNil(req.PostId),
			UserID: uuid.FromStringOrNil(req.UserId),
		}
		if err := tx.Create(&like).Error; err != nil {
			return err
		}
		err := tx.Model(&post).Update("like_count", gorm.Expr("like_count + 1")).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	post.LikeCount++
	return post.ToProto(), nil
}

func (s *Store) HasUserLikedPost(ctx context.Context, postId string, userId string) (bool, error) {
	var like models.PostLike
	if err := s.db.WithContext(ctx).
		Where("post_id = ? AND user_id = ?", postId, userId).
		First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *Store) UnlikePost(ctx context.Context, req *socialv1.UnlikePostRequest) (*socialv1.Post, error) {
	var post models.Post
	if err := s.db.WithContext(ctx).Where("id = ?", req.PostId).First(&post).Error; err != nil {
		return nil, err
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("post_id = ? AND user_id = ?", req.PostId, req.UserId).Delete(&models.PostLike{}).Error; err != nil {
			return err
		}
		err := tx.Model(&post).Update("like_count", gorm.Expr("like_count - 1")).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	post.LikeCount--
	return post.ToProto(), nil
}
