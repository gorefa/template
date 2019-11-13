/**
* @file   : get
* @descrip: 
* @author : ch-yk
* @create : 2018-09-05 下午1:48
* @email  : commonheart.yk@gmail.com
**/

package user

import (

	. "api_gateway/handler"
	"api_gateway/internal/errno"
	"api_gateway/model"

	"github.com/gin-gonic/gin"

)

// @Summary Get an user by the user identifier
// @Description Get an user by username
// @Tags user
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /user/{username} [get]
func Get(c *gin.Context) {
	//获取某个特定用户
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
