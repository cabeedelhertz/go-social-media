package server

import (
	"social/internal/manager"
	"social/proto/gen/social/v1/socialv1connect"
)

var _ socialv1connect.SocialServiceHandler = &Server{}

type Server struct {
	manager *manager.Manager
}

func NewServer(manager *manager.Manager) *Server {
	return &Server{manager: manager}
}
