package middleware

import (
	"net/http"

	"github.com/gol4ng/httpware/metrics"
	"github.com/gol4ng/httpware/middleware"
)

func Metrics(recorder metrics.Recorder) func(http.Handler) http.Handler {
	recorderConfig := metrics.NewConfig(recorder)
	recorderConfig.HandlerNameExtractor = func(req *http.Request) string {
		// Caution: you should keep a few number of unique value
		// it allow you to group metrics measurements
		return req.URL.Path
	}

	return middleware.Metrics(recorderConfig)
}
