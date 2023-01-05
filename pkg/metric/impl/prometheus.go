package impl

import (
	"cep-retriever/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"
)

type Service struct {
	httpRequestHistogram *prometheus.HistogramVec
}

var (
	http = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})
)

func NewPrometheusService() (*Service, error) {
	s := &Service{
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

func (s *Service) SaveHTTP(h *metric.HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
