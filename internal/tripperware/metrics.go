package tripperware

import (
	"net/http"

	"github.com/gol4ng/httpware"
	"github.com/gol4ng/httpware/metrics"
	"github.com/gol4ng/httpware/tripperware"
)

// Configure metrics tripperware
func Metrics(name string, recorder metrics.Recorder) httpware.Tripperware {
	recorderConfig := metrics.NewConfig(recorder)
	recorderConfig.HandlerNameExtractor = func(req *http.Request) string {
		return name
	}
	return tripperware.Metrics(recorderConfig)
}
