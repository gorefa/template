package user

import (
	"fmt"

	. "gogin/handler"
	"gogin/pkg/errno"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	fmt.Println("req", req)
	username := c.Param("username")
	log.Infof("URL username: %s", username)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", req.Username, req.Password)
	if req.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if req.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{
		Username: req.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
