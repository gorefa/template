package main

import (
	"errors"
	"net/http"
	"time"

	"gogin/config"
	"gogin/pkg/logger"
	"gogin/router"
	"gogin/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//gin have release, debug, test mode
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middleware.Logging(),
	)

	// make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logger.L().Fatal("The router has no response, or it might took too long to start up.", logger.Error(err))
		}
		logger.L().Info("The router has been deployed successfully.")
	}()

	logger.L().Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	logger.L().Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/api/v1/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logger.L().Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
