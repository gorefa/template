package k8s

import (
	. "gogin/handler"
	"gogin/model"
	"gogin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func NodeList(c *gin.Context) {
	clusterName := c.Query("cluster")
	if clusterName == "" {
		SendResponse(c, errno.ErrClusterNotSpecified, nil)
		return
	}
	nodelist := model.NodeList(clusterName)
	SendResponse(c, nil, nodelist)
}
