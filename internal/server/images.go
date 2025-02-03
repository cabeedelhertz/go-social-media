package server

import (
	"context"
	socialv1 "social/proto/gen/social/v1"

	"connectrpc.com/connect"
)

func (s *Server) AddImage(ctx context.Context, req *connect.Request[socialv1.AddImageRequest]) (*connect.Response[socialv1.AddImageResponse], error) {
	resp, err := s.manager.AddImage(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(resp), nil
}
