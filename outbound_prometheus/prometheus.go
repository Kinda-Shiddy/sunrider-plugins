//outbound_prometheus/prometheus.go

package outbound_prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterPrometheusIntegration(metrics *PrometheusMetrics) {
	prometheus.MustRegister(metrics.linkCheckStatus)
	prometheus.MustRegister(metrics.linkCheckDuration)
}

type PrometheusMetrics struct {
	linkCheckStatus   *prometheus.GaugeVec
	linkCheckDuration prometheus.Histogram
}

func NewPrometheusMetrics() *PrometheusMetrics {
	histogramOpts := prometheus.HistogramOpts{
		Name:    "link_check_duration_seconds",
		Help:    "Duration of the link checks",
		Buckets: []float64{0.1, 0.5, 1, 2, 5, 10},
	}

	return &PrometheusMetrics{
		linkCheckStatus: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "link_check_status",
			Help: "The status of the link check",
		}, []string{"link", "status_code"}),
		linkCheckDuration: prometheus.NewHistogram(histogramOpts),
	}
}

func RegisterBlackboxExporterMetrics(metrics *PrometheusMetrics) {
	RegisterPrometheusIntegration(metrics)
}
