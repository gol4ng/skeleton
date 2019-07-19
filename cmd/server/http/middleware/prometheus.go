package middleware

import (
	"github.com/gorilla/mux"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"net/http"
)

func Prometheus() mux.MiddlewareFunc {
	httpMetricsMiddleware := middleware.New(middleware.Config{
		GroupedStatus: true,
		Recorder:      prometheus.NewRecorder(prometheus.Config{}),
	})
	// Caution for next middleware if you have a lot of route or dynamic url (that count of unique route)
	// you should put this middleware over each handler you want to monitor
	return func(h http.Handler) http.Handler { return httpMetricsMiddleware.Handler("", h) }
}
