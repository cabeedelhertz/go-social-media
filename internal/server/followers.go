package server

import (
	"context"
	socialv1 "social/proto/gen/social/v1"

	"connectrpc.com/connect"
)

func (s *Server) FollowUser(ctx context.Context, req *connect.Request[socialv1.FollowUserRequest]) (*connect.Response[socialv1.FollowUserResponse], error) {
	resp, err := s.manager.FollowUser(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) UnfollowUser(ctx context.Context, req *connect.Request[socialv1.UnfollowUserRequest]) (*connect.Response[socialv1.UnfollowUserResponse], error) {
	resp, err := s.manager.UnfollowUser(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) ListFollowers(ctx context.Context, req *connect.Request[socialv1.ListFollowersRequest]) (*connect.Response[socialv1.ListFollowersResponse], error) {
	resp, err := s.manager.ListFollowers(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) IsFollowing(ctx context.Context, req *connect.Request[socialv1.IsFollowingRequest]) (*connect.Response[socialv1.IsFollowingResponse], error) {
	isFollowing, err := s.manager.GetIsFollowing(ctx, req.Msg.GetUserId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&socialv1.IsFollowingResponse{
		Following: isFollowing,
	}), nil
}
