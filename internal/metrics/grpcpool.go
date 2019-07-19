package metrics

import (
	"time"

	"github.com/processout/grpc-go-pool"
)

type Pool struct {
	*grpcpool.Pool

	name     string
	recorder *Recorder

	stopCollectMetrics chan struct{}
}

func (p *Pool) GetName() string {
	return p.name
}

func (p *Pool) Close() {
	close(p.stopCollectMetrics)
	p.Pool.Close()
}

func (p *Pool) CollectMetrics(interval *time.Ticker) {
	p.stopCollectMetrics = make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-interval.C:
				p.recorder.CollectGrpcPoolState(p)
			case <-p.stopCollectMetrics:
				interval.Stop()
				return
			}
		}
	}()
}

func NewGrpcPool(name string, pool *grpcpool.Pool, rec *Recorder) *Pool {
	return &Pool{
		name:     name,
		Pool:     pool,
		recorder: rec,
	}
}
