package k8s

import (
	. "gogin/handler"
	"gogin/model"
	"gogin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func PodList(c *gin.Context) {

	clusterName := c.Query("cluster")
	if clusterName == "" {
		SendResponse(c, errno.ErrClusterNotSpecified, nil)
		return
	}
	list := model.PodList(clusterName,c.Param("ns"))
	SendResponse(c, nil, list)
}
