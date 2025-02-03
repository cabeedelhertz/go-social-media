package models

import (
	socialv1 "social/proto/gen/social/v1"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

/*

CREATE TABLE IF NOT EXISTS followers (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid NOT NULL REFERENCES users(id),
    follower_id uuid NOT NULL REFERENCES users(id),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);
*/

type Follower struct {
	ID         uuid.UUID      `gorm:"id"`
	UserID     uuid.UUID      `gorm:"user_id"`
	FollowerID uuid.UUID      `gorm:"follower_id"`
	CreatedAt  time.Time      `gorm:"created_at"`
	UpdatedAt  time.Time      `gorm:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"deleted_at"`
}

func (f *Follower) ToProto() *socialv1.Follower {
	return &socialv1.Follower{
		Id:         f.ID.String(),
		UserId:     f.UserID.String(),
		FollowerId: f.FollowerID.String(),
		CreatedAt:  timestamppb.New(f.CreatedAt),
		UpdatedAt:  timestamppb.New(f.UpdatedAt),
	}
}
