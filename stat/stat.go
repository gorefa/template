package stat

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GaugeVecApiMethod = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiCount",
		Help: "各种网络请求次数",
	}, []string{"method"})

	GaugeVecApiDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiDuration",
		Help: "api耗时单位ms",
	}, []string{"path"})
	GaugeVecApiError = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiErrorCount",
		Help: "请求api错误的次数type: api/ws",
	}, []string{"type"})

	SummaryHttpRequest = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:	"http_request_summary_seconds",
		Help:	"Summary of lantencies for HTTP requests",
	})
	HistogramHttpRequest = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:		"http_request_duration_seconds",
		Help:		"Histogram of lantencies for HTTP requests",
	})
)

func init() {
	prometheus.MustRegister(GaugeVecApiMethod, GaugeVecApiDuration, GaugeVecApiError,HistogramHttpRequest,SummaryHttpRequest)
}
