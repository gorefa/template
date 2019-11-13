/**
* @file   : router.go
* @descrip: 总路由入口 (把相应的请求交给相关的 handler), handler 一口气吃不完的，才会去调用 service 包解决
* @author : ch-yk
* @create : 2018-08-30 下午12:47
* @email  : commonheart.yk@gmail.com
* ****************************************
* change activity:
*  目前只负责 xserver_status 分组
**/

package router

import (
	"net/http"
	"gogin/handler/user"

	"github.com/gin-gonic/gin"
	"gogin/handler/xserver_status"
	"gogin/router/middleware"
)

func Load(engine *gin.Engine, middles ...gin.HandlerFunc) *gin.Engine {
	//gin 的 USE 方法主要利用中间件处理请求头

	//处理可能存在的异常情况(错误，一般是error, panic)导致的服务器服务继续进行下次服务，此时立即恢复
	engine.Use(gin.Recovery())

	//一般中间件都需要处理响应头
	engine.Use(middleware.NoCache)
	engine.Use(middleware.Options)
	engine.Use(middleware.Secure)

	// 设置各类handler
	engine.Use(middles...)


	//注册: 处理404
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "API 路由错误(未定义) 404.")
	})

	/*根目录: TODO: 用模板显示当前 API 能够提供的所有服务 (非当前重点)*/
	engine.Group("/")
	engine.GET("/", func (context *gin.Context) {
		message := "<!DOCTYPE html><html><head><link rel=\"shortcut icon\" href=\"https://hub.commonheart-yk.com/wiki/pics/2018/09/17-15-05-sitelogo.png\" />" +
			"<title>YK API 网关</title><style>body {width: 35em;margin: 0 auto;font-family: Tahoma, Verdana, Arial, sans-serif;}</style></head><body>" +
			"<br/><br/><br/><br/><br/><br/><br/><h1 align=\"center\">YK CGI 内部网关系统</h1><hr/><h2 align=\"center\">目前内测中...</h2><h3 align=\"center\">" +
			"暂时不对外开放</h3><br/><br/><div align=\"center\"><p>The server is powered by <a href=\"https://gitlab.commonheart-yk.com/ch-yk\"> Y K </a>.</p>" +
			"<p><em>faithfully yours.</em></p></div></body></html>"
		context.Writer.WriteString(message)
		//.String(http.StatusOK, message)
	})

	//** TODO: 分组信息应该从数据库中读取出来 (非当前重点) --- 当前是手动添加路由信息

	/*
	** 非常好记，分组名就是 handler 包名 + `_group`,
	*/
	//服务器状态检查组
	xserver_status_group := engine.Group("/v1/xserver_status")
	xserver_status_group.GET("/alive", xserver_status.AliveCheck)
	xserver_status_group.GET("/disk", xserver_status.DiskCheck)
	xserver_status_group.GET("/cpu", xserver_status.CPUCheck)
	xserver_status_group.GET("/mem", xserver_status.MEMCheck)


	/********* 新增 user 模块 (为 '成就系统') ************/
	user_group := engine.Group("/v1/user")
	user_group.Use(middleware.AuthMiddleware()) // 改组操作都需要权限

	//新增
	user_group.POST("", user.Create)
	//获取列表
	user_group.GET("", user.List)
	//删除
	user_group.DELETE("/:id", user.Delete)
	//更新
	user_group.PUT("/:id", user.Update)
	//查看某个具体的 user (根据 username)
	user_group.GET("/:username", user.Get)


	//单独设置 login
	// api for authentication functionalities
	engine.POST("/login", user.Login)








	return engine
}
