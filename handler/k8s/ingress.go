package k8s

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/gin-gonic/gin"
)

func IngressList(c *gin.Context) {

	list := model.IngressList(c.Param("ns"))
	SendResponse(c, nil, list)
}
