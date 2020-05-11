package middleware

import (
	"gogin/stat"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func MwPrometheusHttp(c *gin.Context) {
	// Method
	start := time.Now()
	method := c.Request.Method
	path := c.Request.URL.String()
	stat.GaugeVecApiMethod.WithLabelValues(method).Inc()


	c.Next()
	code := c.Writer.Status()
	end := time.Now()
	d := end.Sub(start) / time.Millisecond
	stat.GaugeVecApiDuration.WithLabelValues(path).Set(float64(d))

	// request
	stat.HistogramHttpRequest.WithLabelValues(path,strconv.Itoa(code)).Observe(time.Since(start).Seconds())
	stat.SummaryHttpRequest.Observe(time.Since(start).Seconds())

}
