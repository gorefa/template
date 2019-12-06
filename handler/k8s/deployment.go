package k8s

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

func DeploymentCreate(c *gin.Context) {

	deployment, err := model.DeploymentCreate(c.Param("ns"))
	if err != nil {
		log.Fatal("deployment create failed", err)
	}

	SendResponse(c, nil, deployment)
}
