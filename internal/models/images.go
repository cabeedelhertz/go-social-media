package models

import (
	socialv1 "social/proto/gen/social/v1"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type PostImage struct {
	ID         uuid.UUID      `gorm:"id"`
	PostID     uuid.UUID      `gorm:"post_id"`
	ImageURL   string         `gorm:"image_url"`
	ObjectKey  string         `gorm:"object_key"`
	OrderIndex int            `gorm:"order_index"`
	CreatedAt  time.Time      `gorm:"created_at"`
	UpdatedAt  time.Time      `gorm:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"deleted_at"`
}

func (pi *PostImage) ToProto() *socialv1.PostImage {
	return &socialv1.PostImage{
		Id:        pi.ID.String(),
		PostId:    pi.PostID.String(),
		Url:       pi.ImageURL,
		ObjectKey: pi.ObjectKey,
	}
}

func PostImageFromProto(pi *socialv1.PostImage) *PostImage {
	return &PostImage{
		ID:        uuid.FromStringOrNil(pi.Id),
		PostID:    uuid.FromStringOrNil(pi.PostId),
		ImageURL:  pi.Url,
		ObjectKey: pi.ObjectKey,
	}
}
