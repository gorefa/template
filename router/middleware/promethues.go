package middleware

import (
	"gogin/stat"
	"time"

	"github.com/gin-gonic/gin"
)

func MwPrometheusHttp(c *gin.Context) {
	start := time.Now()
	method := c.Request.Method
	stat.GaugeVecApiMethod.WithLabelValues(method).Inc()

	c.Next()
	// after request
	end := time.Now()
	d := end.Sub(start) / time.Millisecond
	stat.GaugeVecApiDuration.WithLabelValues(method).Set(float64(d))
}
