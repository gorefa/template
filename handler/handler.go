/**
* @file   : handler
* @descrip: 通用handler, 解析客户端参数不统一，但是回写 json, 统一的写 sendResponse()
* @author : ch-yk
* @create : 2018-09-03 下午6:43
* @email  : commonheart.yk@gmail.com
**/

package handler


import (
	"net/http"

	"gogin/internal/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	//无参数的情况
	if code == errno.BindErr.Code {
		c.JSON(http.StatusBadRequest, Response{
			Code: code,
			Message: message,
			Data: nil,
		}) //gin.H{"error": errno.BindErr}
	}

	// 有参数的情况
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
