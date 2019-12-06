package k8s

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/gin-gonic/gin"
)

func NodeList(c *gin.Context) {

	nodelist := model.NodeList()
	SendResponse(c, nil, nodelist)
}
