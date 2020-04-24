package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/gorefa/log"
	. "gogin/handler"
	"gogin/model"
)

func ClusterCreate(c *gin.Context) {
	cluster := model.Cluster{Enable: true}
	if err := c.ShouldBindJSON(&cluster); err != nil {
		log.Error("bind error",err)

		return
	}
	if err := model.ClusterCreate(cluster); err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil)
}

