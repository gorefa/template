package k8s

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/gin-gonic/gin"
)

func PodList(c *gin.Context) {

	list := model.PodList(c.Param("ns"))
	SendResponse(c, nil, list)
}
