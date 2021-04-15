package router

import (
	"net/http"

	"gogin/handler"
	"gogin/router/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	//g.Use(middleware.OpenTracing())
	g.Use(middleware.TraceSpan())
	g.Use(middleware.RequestIDHandler())
	g.Use(mw...)

	g.POST("/register", handler.Register) // TODO
	g.POST("/login", handler.Login)
	g.GET("/verify", handler.Verify)
	g.GET("/refresh", handler.Refresh)
	g.GET("/sayHello", handler.SayHello)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := g.Group("/api/v1/").Use(middleware.MwPrometheusHttp)
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/disk", handler.DiskCheck)
		v1.GET("/cpu", handler.CPUCheck)
		v1.GET("/ram", handler.RAMCheck)
		v1.PUT("/test", handler.Test)
		v1.POST("/time", handler.Time)
	}

	return g
}
