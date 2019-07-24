package config

import (
	"context"

	"github.com/gol4ng/skeleton/pkg/config"
)

// this config is related to cmd/server
type Server struct {
	ServerHTTPAddr string `config:"http_addr"`
	ServerGRPCPort string `config:"grpc_port"`
}

func NewServer() *Config {
	cfg := &Config{}

	config.LoadOrFatal(context.Background(), cfg)
	// @TODO validate config here
	return cfg
}

// Validate Server config in internal ?
func ValidateServer() {
}
