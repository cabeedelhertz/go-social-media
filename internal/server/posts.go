package server

import (
	"context"
	socialv1 "social/proto/gen/social/v1"

	"connectrpc.com/connect"
)

func (s *Server) LikePost(ctx context.Context, req *connect.Request[socialv1.LikePostRequest]) (*connect.Response[socialv1.LikePostResponse], error) {
	resp, err := s.manager.LikePost(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) UnlikePost(ctx context.Context, req *connect.Request[socialv1.UnlikePostRequest]) (*connect.Response[socialv1.UnlikePostResponse], error) {
	resp, err := s.manager.UnlikePost(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) HasUserLikedPost(ctx context.Context, req *connect.Request[socialv1.HasUserLikedPostRequest]) (*connect.Response[socialv1.HasUserLikedPostResponse], error) {
	hasLiked, err := s.manager.HasUserLikedPost(ctx, req.Msg.PostId, req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&socialv1.HasUserLikedPostResponse{Liked: hasLiked}), nil
}
