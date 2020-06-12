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
		Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2},
	},
	[]string{"job_type"},
)

var jobCounter200 = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "http_requests_total",
		ConstLabels: prometheus.Labels{"response_code": "200"},
	},
)

var jobCounter500 = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name:        "http_requests_total",
		Help:        "http_requests_total",
		ConstLabels: prometheus.Labels{"response_code": "500"},
	},
)

var isLong = true

func init() {
	prometheus.MustRegister(jobsDurationHistogram)
	prometheus.MustRegister(jobCounter200)
	prometheus.MustRegister(jobCounter500)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status/{status:[0-9]+}", status).Methods("GET", "HEAD")

	http.Handle("/", rtr)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func status(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	params := mux.Vars(r)
	status, _ := strconv.Atoi(params["status"])

	if isLong {
		time.Sleep(20 * time.Millisecond)
	}
	isLong = !isLong
	w.WriteHeader(status)
	w.Write([]byte(params["status"]))
	duration := time.Since(start)
	jobsDurationHistogram.WithLabelValues("").Observe(duration.Seconds())

	incrementJobCounter(status)
}

func incrementJobCounter(status int) {
	if status >= 500 {
		jobCounter500.Inc()
	} else {
		jobCounter200.Inc()
	}
}
