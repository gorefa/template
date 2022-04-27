package middleware

import (
	"gogin/pkg/logger"

	"gogin/handler"
	"gogin/pkg/errno"
	"gogin/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if _, err := token.ParseToken(header); err != nil {
			logger.L().Error("Parsetoken is error", logger.Error(err))
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
