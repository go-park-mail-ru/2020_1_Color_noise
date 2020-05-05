package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	hits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hits",
		Help: "Number of all hits.",
	})

	er = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "errors_hits",
		Help: "Number of all errors.",
	})

	rps = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "work_time",
	}, []string{"method", "path"})

	errors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors_vector",
	}, []string{"status_code", "path"})

)

func Register()  {
	prometheus.MustRegister(hits)
	prometheus.MustRegister(rps)
	prometheus.MustRegister(er)

	prometheus.MustRegister(errors)
}

func Increase()  {
	hits.Add(1)
}

func IncreaseError()  {
	er.Add(1)
}

func WorkTime(method, path string, time time.Duration)  {
	rps.WithLabelValues(method, path).Observe(float64(time))
}

func Errors(st, path string)  {
	errors.WithLabelValues(st, path).Inc()
}
