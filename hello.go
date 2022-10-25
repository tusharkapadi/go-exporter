package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			var curr_no = rand.Intn(10)
			if curr_no < 3 {
				jobsInQueue.WithLabelValues("process").Dec()
				jobsInQueue.WithLabelValues("process").Dec()
				jobsInQueue.WithLabelValues("process").Dec()
			} else if curr_no < 6 {
				jobsInQueue.WithLabelValues("access").Dec()
				jobsInQueue.WithLabelValues("access").Dec()
				jobsInQueue.WithLabelValues("access").Dec()
			}
			jobsInQueue.WithLabelValues("process").Inc()
			jobsInQueue.WithLabelValues("access").Inc()
			time.Sleep(2 * time.Second)

		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_cust_processed_ops_total",
		Help: "The total number of processed events",
	})
)

var jobsInQueue = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "go_cust_jobs_in_queue",
	Help: "Current number of jobs in the queue",
},
	[]string{"job_type"},
)

func init() {
	prometheus.MustRegister(jobsInQueue)
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
