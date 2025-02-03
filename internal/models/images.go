package models

import (
	socialv1 "social/proto/gen/social/v1"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

/*
CREATE TABLE IF NOT EXISTS post_images (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    post_id uuid NOT NULL REFERENCES posts(id),
    image_url varchar NOT NULL,
    object_key varchar NOT NULL,
    order_index int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);
*/

type PostImage struct {
	ID     uuid.UUID `gorm:"id"`
	PostID uuid.UUID `gorm:"post_id"`
	// Post       Post
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
		// OrderIndex: pi.OrderIndex,
		// CreatedAt:  timestamppb.New(pi.CreatedAt),
		// UpdatedAt:  timestamppb.New(pi.UpdatedAt),
	}
}

func PostImageFromProto(pi *socialv1.PostImage) *PostImage {
	return &PostImage{
		ID:        uuid.FromStringOrNil(pi.Id),
		PostID:    uuid.FromStringOrNil(pi.PostId),
		ImageURL:  pi.Url,
		ObjectKey: pi.ObjectKey,
		// OrderIndex: pi.OrderIndex,
		// CreatedAt:  pi.CreatedAt.AsTime(),
		// UpdatedAt:  pi.UpdatedAt.AsTime(),
	}
}
