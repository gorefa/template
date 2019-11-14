package handler

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	message := "OK"
	SendResponse(c, nil, message)
	//c.String(http.StatusOK, message)
}
