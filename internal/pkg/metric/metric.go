package metric

import "github.com/prometheus/client_golang/prometheus"

var (
	hits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hits",
		Help: "Number of all hits.",
	})

	rps = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "rpc",
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