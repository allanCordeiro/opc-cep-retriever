package middlewares

import (
	"cep-retriever/pkg/metric"
	"cep-retriever/pkg/metric/impl"
	"log"
	"net/http"
	"strconv"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status string
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = strconv.Itoa(status)
	r.ResponseWriter.WriteHeader(status)
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricService, err := impl.NewPrometheusService()
		if err != nil {
			log.Fatal(err.Error())
		}
		appMetric := metric.NewHTTP(r.URL.Path, r.Method)
		appMetric.Started()
		sw := &StatusRecorder{
			w,
			"200",
		}
		next.ServeHTTP(sw, r)
		appMetric.Finished()
		appMetric.StatusCode = sw.Status
		metricService.SaveHTTP(appMetric)
	})
}
