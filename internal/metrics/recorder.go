package metrics

import "github.com/prometheus/client_golang/prometheus"

type Recorder struct {
	grpcPoolGaugeVec *prometheus.GaugeVec
}

// It register all the metris on the given registerer
// pass nil will use DefaultRegisterer
func (r *Recorder) MustRegister(registry prometheus.Registerer) *Recorder {
	if registry == nil {
		registry = prometheus.DefaultRegisterer
	}

	registry.MustRegister(
		r.grpcPoolGaugeVec,
	)
	return r
}

func (r *Recorder) CollectGrpcPoolState(pool *Pool) *Recorder {
	name := pool.GetName()
	r.grpcPoolGaugeVec.WithLabelValues(name, "capacity").Set(float64(pool.Capacity()))
	r.grpcPoolGaugeVec.WithLabelValues(name, "available").Set(float64(pool.Available()))
	return r
}

func NewRecorder(namespace string) *Recorder {
	return &Recorder{
		grpcPoolGaugeVec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: "grpc_pool",
				Name:      "connection_count",
				Help:      "This represent the number of grpc connection",
			},
			[]string{"service", "type"},
		),
	}
}
