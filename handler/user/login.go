/**
* @file   : login
* @descrip: 处理用户登录行为
* @author : ch-yk
* @create : 2018-09-05 下午5:53
* @email  : commonheart.yk@gmail.com
**/

package user

import (

	"api_gateway/internal/auth"
	"api_gateway/internal/errno"
	"api_gateway/internal/token"
	"api_gateway/model"
	. "api_gateway/handler"

	"github.com/gin-gonic/gin"
)

// 登录成功产生具体的 token 返回给客户端
func Login(c *gin.Context) {
	// Binding the data with the user struct.
	//username, password
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.BindErr, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
