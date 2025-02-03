package databasedriver

import (
	"context"
	"errors"
	"social/internal/models"
	socialv1 "social/proto/gen/social/v1"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func (s *Store) CreateFollower(ctx context.Context, userId, followerId uuid.UUID) (*socialv1.Follower, error) {
	if userId == uuid.Nil || followerId == uuid.Nil {
		return nil, errors.New("invalid create follower request")
	}
	follower := models.Follower{
		UserID:     userId,
		FollowerID: followerId,
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&follower).Error; err != nil {
			return err
		}
		follower := models.User{
			ID: followerId,
		}
		if err := tx.Model(&follower).Update("following_count", gorm.Expr("following_count + 1")).Error; err != nil {
			return err
		}
		user := models.User{
			ID: userId,
		}
		if err := tx.Model(&user).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return follower.ToProto(), nil
}

func (s *Store) DeleteFollower(ctx context.Context, userId, followerId uuid.UUID) error {
	if userId == uuid.Nil || followerId == uuid.Nil {
		return errors.New("invalid delete follower request")
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ? AND follower_id = ?", userId, followerId).Delete(&models.Follower{}).Error
		if err != nil {
			return err
		}
		follower := models.User{
			ID: followerId,
		}
		if err := tx.Model(&follower).Update("following_count", gorm.Expr("following_count - 1")).Error; err != nil {
			return err
		}
		user := models.User{
			ID: userId,
		}
		if err := tx.Model(&user).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *Store) ListFollowers(ctx context.Context, req *socialv1.ListFollowersRequest) ([]*socialv1.Follower, error) {
	var followers []models.Follower
	query := s.db.WithContext(ctx).Where("user_id = ?", req.UserId)
	if err := query.Find(&followers).Error; err != nil {
		return nil, err
	}
	resp := make([]*socialv1.Follower, 0, len(followers))
	for _, follower := range followers {
		resp = append(resp, follower.ToProto())
	}
	return resp, nil
}

func (s *Store) GetIsFollowing(ctx context.Context, userId, followerId uuid.UUID) (bool, error) {
	var follower models.Follower
	if err := s.db.WithContext(ctx).
		Where("user_id = ? AND follower_id = ?", userId, followerId).
		First(&follower).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
