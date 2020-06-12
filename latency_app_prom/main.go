package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var jobsDurationHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_milliseconds",
		Help:    "http request duration milliseconds",
		Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2, 5},
	},
	[]string{"handler"},
)

func init() {
	prometheus.MustRegister(jobsDurationHistogram)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/latency/{latency:[0-9]+}", latency).Methods("GET", "HEAD")

	http.Handle("/", rtr)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func latency(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	params := mux.Vars(r)
	latency, _ := strconv.Atoi(params["latency"])
	time.Sleep(time.Duration(latency) * time.Millisecond)
	w.WriteHeader(200)
	w.Write([]byte(params["latency"]))
	duration := time.Since(start)
	jobsDurationHistogram.WithLabelValues("/latency").Observe(duration.Seconds())
}
