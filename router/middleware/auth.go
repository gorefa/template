package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gogin/handler"
	"gogin/pkg/errno"
	"gogin/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context){
		header  :=  c.Request.Header.Get("Authorization")
		fmt.Println("headert",header)
		//if len(header) == 0 {
		//	handler.SendResponse(c,errno.ErrHeaderNull,nil)
		//	c.Abort()
		//	return
		//}
		if _,err := token.ParseToken(header); err != nil {
			handler.SendResponse(c,errno.ErrTokenInvalid,nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
