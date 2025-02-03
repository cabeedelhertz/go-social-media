package service

import (
	"social/pkg/config"
	"time"
)

type Config struct {
	config.Base       `config:",squash"`
	ValidIssuers      []string      `config:"valid_issuers"`
	GrpcPort          int           `config:"grpc_port"`
	HttpPort          int           `config:"http_port"`
	ReadHeaderTimeout time.Duration `config:"read_header_timeout"`
	ReadTimeout       time.Duration `config:"read_timeout"`
	WriteTimeout      time.Duration `config:"write_timeout"`
}

func (c Config) Root() config.Base {
	return c.Base
}
