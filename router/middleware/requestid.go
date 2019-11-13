/**
* @file   : requestid
* @descrip: gin 引擎需要加载的全局中间件，获取或者设置request id
* @author : ch-yk
* @create : 2018-09-05 下午4:38
* @email  : commonheart.yk@gmail.com
**/

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

//requestId) 设置在返回包的 Header 中
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4:= uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)//在返回包中设置 uuid
		c.Next()
	}
}
