package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type PostLike struct {
	ID        uuid.UUID      `gorm:"id"`
	PostID    uuid.UUID      `gorm:"post_id"`
	UserID    uuid.UUID      `gorm:"user_id"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}
