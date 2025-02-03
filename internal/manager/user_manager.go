package manager

import (
	"context"
	socialv1 "social/proto/gen/social/v1"
)

func (m *Manager) CreateOrUpdateUser(ctx context.Context, req *socialv1.CreateOrUpdateUserRequest) (*socialv1.CreateOrUpdateUserResponse, error) {
	user, err := m.store.CreateOrUpdateUser(ctx, req.User)
	if err != nil {
		return nil, err
	}
	return &socialv1.CreateOrUpdateUserResponse{User: user}, nil
}

func (m *Manager) GetUser(ctx context.Context, req *socialv1.GetUserRequest) (*socialv1.GetUserResponse, error) {
	user, err := m.store.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &socialv1.GetUserResponse{User: user}, nil
}

func (m *Manager) ListUsers(ctx context.Context, req *socialv1.ListUsersRequest) (*socialv1.ListUsersResponse, error) {
	users, err := m.store.ListUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return &socialv1.ListUsersResponse{Users: users}, nil
}
