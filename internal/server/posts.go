package server

import (
	"context"
	socialv1 "social/proto/gen/social/v1"

	"connectrpc.com/connect"
)

func (s *Server) CreatePost(ctx context.Context, req *connect.Request[socialv1.CreatePostRequest]) (*connect.Response[socialv1.CreatePostResponse], error) {
	resp, err := s.manager.CreatePost(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) UpdatePost(ctx context.Context, req *connect.Request[socialv1.UpdatePostRequest]) (*connect.Response[socialv1.UpdatePostResponse], error) {
	resp, err := s.manager.UpdatePost(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}
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

func (s *Server) PublishPost(ctx context.Context, req *connect.Request[socialv1.PublishPostRequest]) (*connect.Response[socialv1.PublishPostResponse], error) {
	resp, err := s.manager.PublishPost(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) GetPost(ctx context.Context, req *connect.Request[socialv1.GetPostRequest]) (*connect.Response[socialv1.GetPostResponse], error) {
	resp, err := s.manager.GetPost(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) GetPostImage(ctx context.Context, req *connect.Request[socialv1.GetPostImageRequest]) (*connect.Response[socialv1.GetPostImageResponse], error) {
	resp, err := s.manager.GetPostImage(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}

func (s *Server) ListPosts(ctx context.Context, req *connect.Request[socialv1.ListPostsRequest]) (*connect.Response[socialv1.ListPostsResponse], error) {
	resp, err := s.manager.ListPosts(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}
