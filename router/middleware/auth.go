package middleware

import (
	"github.com/gorefa/log"

	"gogin/handler"
	"gogin/pkg/errno"
	"gogin/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if _, err := token.ParseToken(header); err != nil {
			log.Error("Parsetoken is error", err)
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
