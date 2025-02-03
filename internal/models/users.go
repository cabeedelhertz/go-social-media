package models

import (
	socialv1 "social/proto/gen/social/v1"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `gorm:"id"`
	ExternalAuthID string         `gorm:"external_auth_id"`
	FirstName      string         `gorm:"first_name"`
	LastName       string         `gorm:"last_name"`
	Email          string         `gorm:"email"`
	Username       string         `gorm:"username"`
	PhotoURL       string         `gorm:"photo_url"`
	Bio            string         `gorm:"bio"`
	FollowerCount  int            `gorm:"follower_count"`
	FollowingCount int            `gorm:"following_count"`
	LastLoginTime  time.Time      `gorm:"last_login_time"`
	CreatedAt      time.Time      `gorm:"created_at"`
	UpdatedAt      time.Time      `gorm:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"deleted_at"`
}

func (u *User) ToProto() *socialv1.User {
	return &socialv1.User{
		Id:             u.ID.String(),
		ExternalAuthId: u.ExternalAuthID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		Username:       u.Username,
		PhotoUrl:       u.PhotoURL,
		Bio:            u.Bio,
		FollowerCount:  int64(u.FollowerCount),
		FollowingCount: int64(u.FollowingCount),
		LastLoginTime:  timestamppb.New(u.LastLoginTime),
		CreatedAt:      timestamppb.New(u.CreatedAt),
		UpdatedAt:      timestamppb.New(u.UpdatedAt),
	}
}

func UserFromProto(user *socialv1.User) *User {
	return &User{
		ID:             uuid.FromStringOrNil(user.Id),
		ExternalAuthID: user.ExternalAuthId,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Username:       user.Username,
		PhotoURL:       user.PhotoUrl,
		Bio:            user.Bio,
		LastLoginTime:  user.LastLoginTime.AsTime(),
		CreatedAt:      user.CreatedAt.AsTime(),
		UpdatedAt:      user.UpdatedAt.AsTime(),
	}
}
