package user

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	Users []string
}

func List(c *gin.Context) {
	users := model.List()

	rsp := ListResponse{
		Users: users,
	}

	SendResponse(c, nil, rsp)
}
