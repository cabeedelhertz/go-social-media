package service

import (
	"context"
	"errors"
	"fmt"
	"social/pkg/common/logging"
	"social/pkg/service/system"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

const LoggerName = "service.runner"

var DefaultMeterProvider metric.MeterProvider = noop.NewMeterProvider()

type Server interface {
	Start(context.Context) error
	Stop() error
}

type Service struct {
	logger       *zap.Logger
	systemServer *system.Server
	servers      []Server
	cfg          Config
}

func wrap(logger *zap.Logger) func(msg string, args ...interface{}) {
	return func(msg string, args ...interface{}) {
		logger.Info(fmt.Sprintf(msg, args...))
	}
}

func New(cfg Config, servers ...Server) (*Service, error) {

	logger := logging.NewLogger(LoggerName)
	maxprocs.Set(maxprocs.Logger(wrap(logger)))

	server, err := system.NewServer(cfg.Root())
	if err != nil {
		return nil, err
	}

	return &Service{
		logger:       logging.NewLogger(LoggerName),
		cfg:          cfg,
		systemServer: server,
		servers:      servers,
	}, nil
}

func (s *Service) Run(ctx context.Context) error {
	return SignalRunner(logging.NewContextLogger(ctx, s.logger), s.start, s.stop)
}

func (s *Service) start(ctx context.Context) error {

	errs := make(chan error)

	s.run(ctx, s.systemServer, errs)
	for _, server := range s.servers {
		s.run(ctx, server, errs)
	}

	for err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) run(ctx context.Context, server Server, errs chan error) {
	go func() {
		errs <- server.Start(ctx)
	}()
}

func (s *Service) stop() (err error) {
	for _, server := range s.servers {
		if serr := server.Stop(); serr != nil {
			err = errors.Join(err, serr)
		}
	}
	if serr := s.systemServer.Stop(); serr != nil {
		err = errors.Join(err, serr)
	}
	return
}
