package service

import (
	"context"
	"os"
	"os/signal"
	"social/pkg/common/logging"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type RunFunc func(context.Context) error

type StopFunc func() error

func SignalRunner(ctx context.Context, runner RunFunc, shutdown StopFunc) error {

	logger := logging.FromContext(ctx)

	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	complete := make(chan error)
	go func() {
		defer close(complete)
		complete <- runner(ctx)
	}()

	select {
	case sig := <-signals:
		logger.Info("stopping from signal", zap.String("signal", sig.String()))
	case <-ctx.Done():
		logger.Info("context done, stopping")
	case err := <-complete:
		if err != nil {
			logger.Error("failed to start", zap.Error(err))
			return err
		}
		logger.Info("stopping from complete")
	}

	done := make(chan error)
	go func() {
		done <- shutdown()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(20 * time.Second):
		logger.Info("time out, exiting")
	}

	return nil
}
