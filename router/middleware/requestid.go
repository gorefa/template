package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKey int

const ridKey = ctxKey(0)

// NewContext creates a context with request id
func NewContext(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, ridKey, rid)
}

// FromContext returns the request id from context
func FromContext(ctx context.Context) (string) {
	rid  := ctx.Value(ridKey).(string)
	return rid
}

func RequestIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.Request.Header.Get("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
			c.Request.Header.Set("X-Request-ID", rid)
		}
		ctx := NewContext(c.Request.Context(), rid)
		c.Request.WithContext(ctx)
		c.Next()
	}
}
