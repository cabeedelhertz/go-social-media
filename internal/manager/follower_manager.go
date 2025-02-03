package manager

import (
	"context"
	"social/pkg/service/auth"
	socialv1 "social/proto/gen/social/v1"

	"github.com/gofrs/uuid"
)

func (m *Manager) FollowUser(ctx context.Context, req *socialv1.FollowUserRequest) (*socialv1.FollowUserResponse, error) {
	subj := auth.SubjectFromContext(ctx)
	_, err := m.store.CreateFollower(ctx, uuid.FromStringOrNil(req.GetUserId()), uuid.FromStringOrNil(subj.ID))
	if err != nil {
		return nil, err
	}
	return &socialv1.FollowUserResponse{}, nil
}

func (m *Manager) UnfollowUser(ctx context.Context, req *socialv1.UnfollowUserRequest) (*socialv1.UnfollowUserResponse, error) {
	subj := auth.SubjectFromContext(ctx)
	err := m.store.DeleteFollower(ctx, uuid.FromStringOrNil(req.GetUserId()), uuid.FromStringOrNil(subj.ID))
	if err != nil {
		return nil, err
	}
	return &socialv1.UnfollowUserResponse{}, nil
}

func (m *Manager) ListFollowers(ctx context.Context, req *socialv1.ListFollowersRequest) (*socialv1.ListFollowersResponse, error) {
	followers, err := m.store.ListFollowers(ctx, req)
	if err != nil {
		return nil, err
	}
	return &socialv1.ListFollowersResponse{Followers: followers}, nil
}

func (m *Manager) GetIsFollowing(ctx context.Context, userId string) (bool, error) {
	subj := auth.SubjectFromContext(ctx)
	return m.store.GetIsFollowing(ctx, uuid.FromStringOrNil(userId), uuid.FromStringOrNil(subj.ID))
}
