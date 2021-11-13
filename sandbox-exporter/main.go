package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	sandboxGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sandbox_number",
			Help: "sandbox Gauge",
		},
		[]string{"sandbox"},
	)
)

func setRandomValue() {
	for {
		rand.Seed(time.Now().UnixNano())
		n := -1 + rand.Float64()*2
		sandboxGauge.With(prometheus.Labels{"sandbox": "sandbox-exporter-random"}).Set(n)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go setRandomValue()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9111", nil)
}
