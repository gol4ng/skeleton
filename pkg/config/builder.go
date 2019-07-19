package config

import (
	"context"
	"log"
	"os"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"

	"github.com/gol4ng/skeleton/pkg/config/dotenv"
)

var DefaultConfigBuilder = NewDefaultConfigBuilder()

type Builder struct {
	backends []backend.Backend
}

func (cb *Builder) AppendBackends(backends ...backend.Backend) *Builder {
	cb.backends = append(cb.backends, backends...)

	return cb
}

func (cb *Builder) PrependBackends(backends ...backend.Backend) *Builder {
	cb.backends = append(backends, cb.backends...)

	return cb
}

func (cb *Builder) Load(ctx context.Context, to interface{}) error {
	backends := cb.backends[:0]
	for _, b := range cb.backends {
		if b != nil {
			backends = append(backends, b)
		}
	}
	return confita.NewLoader(backends...).Load(ctx, to)
}

func (cb *Builder) LoadOrFatal(ctx context.Context, to interface{}) {
	if err := cb.Load(ctx, to); err != nil {
		log.Fatal(err)
	}
}

func NewConfigBuilder(backends ...backend.Backend) *Builder {
	return &Builder{backends: append([]backend.Backend{}, backends...)}
}

func Load(ctx context.Context, to interface{}) error {
	return DefaultConfigBuilder.Load(ctx, to)
}

func LoadOrFatal(ctx context.Context, to interface{}) {
	DefaultConfigBuilder.LoadOrFatal(ctx, to)
}

/*
 * Create Builder preconfigured with:
 * - .env file loader if file exist
 * - environment variable loader
 * - flags loader
 */
func NewDefaultConfigBuilder() *Builder {
	builder := NewConfigBuilder(
		env.NewBackend(),
		flags.NewBackend(),
	)

	f := ".env"
	if _, err := os.Stat(f); err == nil {
		builder.PrependBackends(dotenv.NewBackend(f))
	}

	return builder
}
