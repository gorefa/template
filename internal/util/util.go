/**
* @file   : util
* @descrip: 
* @author : ch-yk
* @create : 2018-09-05 下午1:54
* @email  : commonheart.yk@gmail.com
**/

package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

/* 拿到每次请求的请求 id */
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}