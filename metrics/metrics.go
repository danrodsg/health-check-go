package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

// Define um Gauge para a latência da checagem de saúde
var HealthCheckLatency = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "health_check_latency_ms",
	Help: "Latency of the full health check operation in milliseconds.",
})

// RecordLatency registra o tempo de execução da checagem.
func RecordLatency(start time.Time) {
	duration := time.Since(start)
	HealthCheckLatency.Set(float64(duration.Milliseconds()))
}

// Outras métricas importantes:
// var TotalRequests = promauto.NewCounter(...)
// var ErrorRequests = promauto.NewCounter(...)