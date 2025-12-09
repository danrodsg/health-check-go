package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)


var HealthCheckLatency = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "health_check_latency_ms",
	Help: "Latency of the full health check operation in milliseconds.",
})


func RecordLatency(start time.Time) {
	duration := time.Since(start)
	HealthCheckLatency.Set(float64(duration.Milliseconds()))
}

