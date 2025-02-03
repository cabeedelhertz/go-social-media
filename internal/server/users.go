package server

import (
	"context"
	socialv1 "social/proto/gen/social/v1"

	"connectrpc.com/connect"
)

func (s *Server) CreateOrUpdateUser(ctx context.Context, req *connect.Request[socialv1.CreateOrUpdateUserRequest]) (*connect.Response[socialv1.CreateOrUpdateUserResponse], error) {
	resp, err := s.manager.CreateOrUpdateUser(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) GetUser(ctx context.Context, req *connect.Request[socialv1.GetUserRequest]) (*connect.Response[socialv1.GetUserResponse], error) {
	resp, err := s.manager.GetUser(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) ListUsers(ctx context.Context, req *connect.Request[socialv1.ListUsersRequest]) (*connect.Response[socialv1.ListUsersResponse], error) {
	resp, err := s.manager.ListUsers(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}
