/**
* @file   : create
* @descrip: 新添加一个 User 对应相关的方法
* @author : ch-yk
* @create : 2018-09-02 下午3:02
* @email  : commonheart.yk@gmail.com
**/

package user

import (

	. "yk_cgi/handler"
	"yk_cgi/internal/errno"
	"yk_cgi/internal/util"
	"yk_cgi/model"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)


// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	// 创建一个新的服务, 错误返回一般流程就是返回 Err，然后 errno.DecodeErr(tmp_err) 结果返回给客户端
	// 期待: POST方法, http://[hostname]/api/v1/user
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	//1. 检查请求操作 (解析客户端传递过来的 json 参数, 并构建后续用于 ORM 的 model)
	//Username string `json:"username"`
	//Password string `json:"password"`
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.BindErr, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		//其他信息不需要，主见 id 自动填充，其他非重要信息字段可以空着
	}
	//检查 UserModel 的参数内容
	//这里借助 v8 框架, Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//2. 执行请求操作
	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	//3. 构建response, 然后写回给客户端
	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}