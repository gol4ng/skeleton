package service

import (
	"github.com/gol4ng/httpware/metrics"
	prom "github.com/gol4ng/httpware/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func (container *Container) GetMetricsRegistry() prometheus.Registerer {
	return prometheus.DefaultRegisterer
}

func (container *Container) GetHttpRecorder() metrics.Recorder {
	if container.httpMetricsRecorder == nil {
		container.httpMetricsRecorder = prom.NewRecorder(prom.Config{})
		container.httpMetricsRecorder.RegisterOn(container.GetMetricsRegistry())
	}
	return container.httpMetricsRecorder
}
