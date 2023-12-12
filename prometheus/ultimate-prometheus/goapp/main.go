package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var criticismCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "aaa_criticism",
	Help: "Total number of times Jackie criticized Dani",
})

var liesCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "aaa_lies",
	Help: "Total number of lies told by politicians",
}, []string{"name"})

var temperatureGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "aaa_temperature",
	Help: "Temperature in celsius",
})

func main() {
	criticismCounter.Inc()
	liesCounter.With(map[string]string{"name": "p1"}).Inc()
	liesCounter.With(map[string]string{"name": "p2"}).Inc()
	liesCounter.With(map[string]string{"name": "p2"}).Inc()

	temperatureGauge.Set(24.4)
	fmt.Println("running")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":1234", nil))

}
