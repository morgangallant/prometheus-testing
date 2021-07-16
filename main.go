package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "testing_requests",
		Help: "The number of requests.",
	})
)

func handler(w http.ResponseWriter, r *http.Request) {
	requests.Inc()
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
