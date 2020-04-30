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

	rps = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "rpc",
	}, []string{"status", "path"})

	rpss = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "rps_summary",
	}, []string{"status", "path"})

)

func Register()  {
	prometheus.MustRegister(hits)
	prometheus.MustRegister(rps)
}

func Increase()  {
	hits.Add(1)
}

func IncreaseRps(status, path string)  {
	rps.WithLabelValues(status, path).Add(1)
}

func WorkTime(method, path string, time time.Duration)  {
	rpss.WithLabelValues(method, path).Observe(float64(time))
}