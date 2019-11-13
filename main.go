package main

import (
	"errors"
	"net/http"
	"time"

	"gogin/config"
	"gogin/model"
	"gogin/router"
	"gogin/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "",
		"./可执行文件 --config 配置文件路径\n或者\n./可执行文件 -c 配置文件路径")
)

/*http: 8076, https: 9076*/
func main() {

	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//初始化数据库连接, 参考 model/db.go
	model.DB.Init()
	defer model.DB.Close()

	/*********** 读取配置 end *************/

	//先设置 gin server 引擎的运行模式 (gin 自带 release, dubug, test 模式)
	gin.SetMode(viper.GetString("run_mode"))

	//拿到 gin server 引擎
	ginEngine := gin.New()
	//ginEngine.Use(middleware.AuthMiddleware()) //全局中间件, 必须认证才能操作(放在路由 Load 里面)

	middleware := []gin.HandlerFunc{
		//middleware.AuthMiddleware(), //放在路由组中，随需要再添加
		middleware.Logging(),
		middleware.RequestId(),
	}

	//开个协程看看路由是否能正常工作
	go func() {
		if err := pingXServer(); err != nil {
			log.Fatal("路由无响应或者相应时间过长.", err)
		}
		log.Info("路由服务已经部署成功")
	}()

	//引入路由，然后由路由交给相应的分组，Load方法，正式引入 api 的 handler --> service
	router.Load(ginEngine, middleware...)

	/*http*/
	go func() {
		log.Infof("开始监听客户端请求, 监听端口: %s:%s", viper.GetString("http_url"), viper.GetString("http_port"))
		log.Info(http.ListenAndServe(":"+viper.GetString("http_port"), ginEngine).Error())
	}()

	/*https*/
	log.Infof("开始监听客户端请求, 监听端口: %s:%s", viper.GetString("https_url"), viper.GetString("https_port"))
	log.Info(http.ListenAndServeTLS(":"+viper.GetString("https_port"),
		viper.GetString("cert.pem"), viper.GetString("cert.key"), ginEngine).Error())

}

/*自检路由是否正常*/
func pingXServer() error {
	//一共 ping 5 次  --- 可以设置指定次数
	for i := 0; i < viper.GetInt("max_ping_tims"); i++ {
		// 发送 `/alive` 来测试是否存活
		resp, err := http.Get(viper.GetString("http_url") + ":" +
			viper.GetString("http_port") + "/v1/xserver_status/alive")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("正在等待路由反馈，1 秒后重试")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到服务器")
}
