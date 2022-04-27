package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"gogin/handler"
	"gogin/pkg/errno"
	"gogin/pkg/logger"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		if path == "/api/v1/health" || path == "/api/v1/disk" || path == "/api/v1/cpu" || path == "/api/v1/ram" {
			return
		}

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""
		ContentType := blw.Header().Get("Content-Type")
		if ContentType == "application/json; charset=utf-8" {
			// get code and message
			var response handler.Response
			if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
				code = errno.InternalServerError.Code
				message = err.Error()
			} else {
				code = response.Code
				message = response.Message
			}

			logger.L().Infof("%-13s | %-12s  | %s %s | {code: %d, message: %s}", latency, ip, method, path, code, message)
		} else {
			logger.L().Infof("%-13s | %-12s | %s %s | ContentType:%s ", latency, ip, method, path, ContentType)

		}

	}
}
