package outbound_prometheus

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterPrometheusIntegration() {
	http.Handle("/metrics", promhttp.Handler())
}

type PrometheusMetrics struct {
	linkCheckStatus   *prometheus.GaugeVec
	linkCheckDuration prometheus.Histogram
}

func NewPrometheusMetrics() *PrometheusMetrics {
	return &PrometheusMetrics{
		linkCheckStatus: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "link_check_status",
			Help: "The status of the link check",
		}, []string{"link", "status_code"}),
		linkCheckDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "link_check_duration_seconds",
			Help: "Duration of the link checks",
		}),
	}
}

func (pm *PrometheusMetrics) SetLinkCheckStatus(link string, statusCode int, value float64) {
	pm.linkCheckStatus.WithLabelValues(link, strconv.Itoa(statusCode)).Set(value)
}

func (pm *PrometheusMetrics) ObserveLinkCheckDuration(duration float64) {
	pm.linkCheckDuration.Observe(duration)
}
