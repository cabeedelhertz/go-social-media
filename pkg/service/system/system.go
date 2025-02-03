package system

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"social/pkg/common/config"
	"social/pkg/common/logging"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	logger     *zap.Logger
	httpServer *http.Server

	cfg config.Base
	mu  sync.RWMutex
}

func NewServer(cfg config.Base) (*Server, error) {
	logger := logging.NewLogger("service.system")

	mux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.SystemPort),
		Handler: mux,
	}

	s := &Server{
		cfg:        cfg,
		logger:     logger,
		httpServer: httpServer,
	}

	return s, nil
}

func (s *Server) RegisterHandler(path string, handler http.Handler) {
	if mux, ok := s.httpServer.Handler.(*http.ServeMux); ok {
		mux.Handle(path, handler)
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting system server", zap.Int("port", s.cfg.SystemPort))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.BindAddress, s.cfg.SystemPort))
	if err != nil {
		return err
	}
	return s.httpServer.Serve(lis)
}

func (s *Server) Stop() error {
	defer s.logger.Sync()
	return s.httpServer.Shutdown(context.Background())
}
