package server

import (
	"context"
	"social/internal/manager"
	socialv1 "social/proto/gen/social/v1"
	"social/proto/gen/social/v1/socialv1connect"

	"connectrpc.com/connect"
)

var _ socialv1connect.SocialServiceHandler = &Server{}

type Server struct {
	manager *manager.Manager
}

func NewServer(manager *manager.Manager) *Server {
	return &Server{manager: manager}
}

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

func (s *Server) AddImage(ctx context.Context, req *connect.Request[socialv1.AddImageRequest]) (*connect.Response[socialv1.AddImageResponse], error) {
	resp, err := s.manager.AddImage(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
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
