package logging

import (
	"context"
	"social/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zcfg zap.Config

func Configure(cfg config.Base) error {
	zcfg = zap.NewDevelopmentConfig()

	if cfg.LogLevel != "" {
		var level zapcore.Level
		if err := level.Set(cfg.LogLevel); err == nil {
			zcfg.Level.SetLevel(level)
		}
	}

	logger, err := zcfg.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)

	return nil
}

func NewLogger(name string) *zap.Logger {
	return zap.L().Named(name)
}

func With(fields ...zap.Field) *zap.Logger {
	return zap.L().With(fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zap.L().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zap.L().Fatal(msg, fields...)
}

type key int

const (
	loggerKey key = iota
)

func NewContextLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	if t, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return t
	}
	return zap.L()
}
