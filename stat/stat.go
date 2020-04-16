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
	}, []string{"WSorAPI"})
)

func init() {
	prometheus.MustRegister(GaugeVecApiMethod, GaugeVecApiDuration)
}
