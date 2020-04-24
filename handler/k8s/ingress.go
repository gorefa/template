package k8s

import (
	. "gogin/handler"
	"gogin/model"
	"gogin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func IngressList(c *gin.Context) {

	clusterName := c.Query("cluster")
	if clusterName == "" {
		SendResponse(c, errno.ErrClusterNotSpecified, nil)
		return
	}
	list := model.IngressList(clusterName,c.Param("ns"))
	SendResponse(c, nil, list)
}
