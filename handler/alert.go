package handler

import (
	"gogin/model"
	"gogin/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Alert
func Alert(c *gin.Context) {

	alert := model.Alerts{}

	if err := c.ShouldBindJSON(&alert); err != nil {
		SendResponse(c, errno.New(errno.ErrBind, err), nil)
		return
	}

	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + c.Param("key")

	err := model.Alertmessage(&alert, url)
	if err != nil {
		SendResponse(c, errno.New(errno.ErrAlert, err), nil)
		return
	}

	c.String(200, "OK")
}
