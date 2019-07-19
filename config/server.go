package config

import (
	"context"

	"github.com/gol4ng/skeleton/pkg/config"
)

// this config is related to cmd/server
type Server struct {
	PprofEnable bool   `config:"pprof"`
	HTTPAddr    string `config:"http_addr"`
	GRPCPort    string `config:"grpc_port"`
}

func NewServer() *Base {
	cfg := &Base{
		Debug: true,
		//Default value here
		Server: Server{
			HTTPAddr: "localhost:8001",
			GRPCPort: "8002",
		},
	}

	config.LoadOrFatal(context.Background(), cfg)
	return cfg
}
