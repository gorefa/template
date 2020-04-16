package router

import (
	"net/http"

	"gogin/handler"
	"gogin/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	v1 := g.Group("/api/v1/")
	{
		v1.GET("/health", handler.HealthCheck)
	}

	g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return g
}
