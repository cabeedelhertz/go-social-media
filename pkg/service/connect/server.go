package connect

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"social/pkg/common/logging"
	"social/pkg/service"
	"social/pkg/service/auth"
	"strings"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type HandlerFactory[T any] func(T, ...connect.HandlerOption) (string, http.Handler)

type Server struct {
	logger     *zap.Logger
	httpServer *http.Server
	mux        *http.ServeMux
	cfg        service.Config
}

type Service[T any] struct {
	Factory HandlerFactory[T]
	Handler T
}

func (s Service[T]) Register(server *Server, opts ...connect.HandlerOption) string {
	path, route := s.Factory(s.Handler, connect.WithInterceptors(authInterceptor()))
	server.mux.Handle(path, route)
	return path
}

func authInterceptor() connect.Interceptor {
	return connect.UnaryInterceptorFunc(func(uf connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
			subject, err := auth.ParseSubjectFromMetadata(ctx)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			if subject != nil {
				ctx = auth.ContextWithSubject(ctx, subject)
			}
			return uf(ctx, request)
		}
	})

}

func NewServer[T any](cfg service.Config, factory HandlerFactory[T], handler T) *Server {
	s := &Server{
		cfg:    cfg,
		logger: logging.NewLogger("service.connect"),
		mux:    http.NewServeMux(),
	}
	service := &Service[T]{factory, handler}
	path := service.Register(s)
	serviceName := strings.ReplaceAll(path, "/", "")
	reflector := grpcreflect.NewStaticReflector(serviceName)
	s.mux.Handle(grpcreflect.NewHandlerV1(reflector))
	s.mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.BindAddress, cfg.GrpcPort),
		Handler:           h2c.NewHandler(cors.New(cors.Options{}).Handler(s.mux), &http2.Server{}),
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting connect server", zap.Int("port", s.cfg.GrpcPort))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.BindAddress, s.cfg.GrpcPort))
	if err != nil {
		return err
	}
	return s.httpServer.Serve(lis)
}

func (s *Server) Stop() error {
	defer s.logger.Sync()

	return s.httpServer.Shutdown(context.Background())
}
