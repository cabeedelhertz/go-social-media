package models

import (
	socialv1 "social/proto/gen/social/v1"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

/*
CREATE TYPE post_status AS ENUM ('DRAFT', 'PUBLISHED', 'ARCHIVED');

CREATE TABLE IF NOT EXISTS posts (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid NOT NULL REFERENCES users(id),
    description varchar NOT NULL,
    status post_status NOT NULL DEFAULT 'DRAFT',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

*/

type PostStatus string

const (
	PostStatusDraft     PostStatus = "DRAFT"
	PostStatusPublished PostStatus = "PUBLISHED"
	PostStatusArchived  PostStatus = "ARCHIVED"
)

// Post represents a post in the database
type Post struct {
	ID           uuid.UUID `gorm:"id"`
	UserID       uuid.UUID `gorm:"user_id"`
	User         User
	Description  string     `gorm:"description"`
	Status       PostStatus `gorm:"status"`
	Images       []PostImage
	Likes        []PostLike
	Comments     []PostComment
	LikeCount    int            `gorm:"like_count"`
	CommentCount int            `gorm:"comment_count"`
	CreatedAt    time.Time      `gorm:"created_at"`
	UpdatedAt    time.Time      `gorm:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"deleted_at"`
}

func (p *Post) ToProto() *socialv1.Post {
	post := &socialv1.Post{
		Id:        p.ID.String(),
		UserId:    p.UserID.String(),
		Caption:   p.Description,
		Status:    PostStatusToProto(p.Status),
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		LikeCount: int64(p.LikeCount),
	}
	if len(p.Images) > 0 {
		for _, img := range p.Images {
			post.Images = append(post.Images, img.ToProto())
			post.ImageUrls = append(post.ImageUrls, img.ImageURL)
		}
	}
	if p.User.ID != uuid.Nil {
		post.Creator = p.User.ToProto()
	}

	return post
}

func PostFromProto(post *socialv1.Post) *Post {
	p := &Post{
		UserID:      uuid.FromStringOrNil(post.UserId),
		Description: post.Caption,
		Status:      PostStatusFromProto(post.Status),
		CreatedAt:   post.CreatedAt.AsTime(),
		UpdatedAt:   post.UpdatedAt.AsTime(),
		LikeCount:   int(post.LikeCount),
	}
	if post.Id != "" {
		p.ID = uuid.FromStringOrNil(post.Id)
	}
	return p
}

func PostStatusFromProto(status socialv1.PostStatus) PostStatus {
	switch status {
	case socialv1.PostStatus_POST_STATUS_DRAFT:
		return PostStatusDraft
	case socialv1.PostStatus_POST_STATUS_PUBLISHED:
		return PostStatusPublished
	case socialv1.PostStatus_POST_STATUS_ARCHIVED:
		return PostStatusArchived
	default:
		return PostStatusDraft
	}
}

func PostStatusToProto(status PostStatus) socialv1.PostStatus {
	switch status {
	case PostStatusDraft:
		return socialv1.PostStatus_POST_STATUS_DRAFT
	case PostStatusPublished:
		return socialv1.PostStatus_POST_STATUS_PUBLISHED
	case PostStatusArchived:
		return socialv1.PostStatus_POST_STATUS_ARCHIVED
	default:
		return socialv1.PostStatus_POST_STATUS_UNSPECIFIED
	}
}
