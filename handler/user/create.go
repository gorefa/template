package user

import (
	"gogin/model"
	"strings"

	. "gogin/handler"
	"gogin/pkg/errno"

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
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			SendResponse(c, &errno.Errno{Code: 400, Message: "user exist!"}, nil)
			return
		}
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: req.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
