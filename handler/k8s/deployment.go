package k8s

import (
	. "gogin/handler"
	"gogin/model"
	"gogin/pkg/errno"

	"github.com/gorefa/log"

	"github.com/gin-gonic/gin"
)

func DeploymentCreate(c *gin.Context) {

	clusterName := c.Query("cluster")
	if clusterName == "" {
		SendResponse(c, errno.ErrClusterNotSpecified, nil)
		return
	}
	deployment, err := model.DeploymentCreate(clusterName,c.Param("ns"))
	if err != nil {
		log.Fatal("deployment create failed", err)
	}

	SendResponse(c, nil, deployment)
}
