package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type PostComment struct {
	ID     uuid.UUID `gorm:"id"`
	PostID uuid.UUID `gorm:"post_id"`
	// Post      Post
	UserID uuid.UUID `gorm:"user_id"`
	// User      User
	Comment   string         `gorm:"comment"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}
