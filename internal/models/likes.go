package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

/*
CREATE TABLE IF NOT EXISTS post_likes (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    post_id uuid NOT NULL REFERENCES posts(id),
    user_id uuid NOT NULL REFERENCES users(id),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);
*/

type PostLike struct {
	ID     uuid.UUID `gorm:"id"`
	PostID uuid.UUID `gorm:"post_id"`
	// Post      Post
	UserID uuid.UUID `gorm:"user_id"`
	// User      User
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}
