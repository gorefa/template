/**
* @file   : create
* @descrip: 新添加一个 Service 对应相关的方法
* @author : ch-yk
* @create : 2018-09-02 下午3:02
* @email  : commonheart.yk@gmail.com
**/

package service


import (
	"fmt"
	"net/http"

	"gogin/internal/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// 创建一个新的服务, 错误返回一般流程就是返回 Err，然后 errno.DecodeErr(tmp_err) 结果返回给客户端
func Create(c *gin.Context) {
	var r struct {
		Servicename string `json:"servicename"`
		Serviceinfo string `json:"serviceinfo"`
	}

	var tmp_err error
	/*检查请求是否传入了 Conent-Type 参数*/
	//如果没有传入参数 && 传递了空 `{}`
	if errinfo := c.Bind(&r); errinfo != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errno.BindErr})
		return
	}

	//如果传入了参数 -- 检查参数，然后构造 Error 用于后台，返回 DecodeErr 信息给 Client
	log.Debugf("servicename is: [%s], serviceinfo is [%s]", r.Servicename, r.Serviceinfo)
	if r.Servicename == "" {
		//fmt.Errorf("后台私密信息, 该服务目前不存在") 用于构造一个标准库的 error, 它对 client 透明
		tmp_err = errno.New(errno.ServiceNotFoundErr, fmt.Errorf("后台私密信息, 该服务目前不存在")).Add("This is add message.")
		log.Errorf(tmp_err, "Get an error")
	}

	if errno.IsServiceNotFoundErr(tmp_err) {
		log.Debug("err type is ServiceNotFoundErr")
	}

	if r.Serviceinfo == "" {
		tmp_err = fmt.Errorf("Serviceinfo is empty")
	}

	code, message := errno.DecodeErr(tmp_err)
	/*返回给客户端的信息*/
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
