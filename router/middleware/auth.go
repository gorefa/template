/**
* @file   : auth.go
* @descrip: 用于权限检查的中间件 (由 gin 引擎加载 Use())
* @author : ch-yk
* @create : 2018-09-05 下午4:32
* @email  : commonheart.yk@gmail.com
**/

package middleware

import (
	"github.com/gin-gonic/gin"
	"yk_cgi/handler"
	"yk_cgi/internal/errno"
	"yk_cgi/internal/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token (同时检查)
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}