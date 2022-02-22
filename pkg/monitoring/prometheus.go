package monitoring

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(router *mux.Router) {
	router.Handle("/metrics", promhttp.Handler())
}

var (
	TotalPersonsCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "person_created",
		Help: "The total number persons created",
	})
)
