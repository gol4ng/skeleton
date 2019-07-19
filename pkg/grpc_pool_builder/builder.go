package grpc_pool_builder

import (
	"time"

	"google.golang.org/grpc"

	"github.com/processout/grpc-go-pool"
)

type PoolBuilder struct {
	dialOpts            []grpc.DialOption
	poolMaxLifeDuration []time.Duration
}

func (b PoolBuilder) DialOpts(opts ...grpc.DialOption) PoolBuilder {
	b.dialOpts = append(b.dialOpts, opts...)
	return b
}

func (b PoolBuilder) PoolMaxLifeDuration(maxLifeDuration ...time.Duration) PoolBuilder {
	b.poolMaxLifeDuration = append(b.poolMaxLifeDuration, maxLifeDuration...)
	return b
}

func (b PoolBuilder) Build(url string, init int, capacity int, idleTimeout time.Duration) (*grpcpool.Pool, error) {
	return b.getPool(url, init, capacity, idleTimeout)
}

func (b PoolBuilder) getPool(url string, init int, capacity int, idleTimeout time.Duration) (*grpcpool.Pool, error) {
	return grpcpool.New(b.getFactory(url), init, capacity, idleTimeout, b.poolMaxLifeDuration...)
}

func (b PoolBuilder) getFactory(url string) grpcpool.Factory {
	return func() (*grpc.ClientConn, error) {
		return grpc.Dial(url, b.dialOpts...)
	}
}
