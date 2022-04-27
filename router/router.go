package router

import (
	"net/http"

	"gogin/handler"
	"gogin/router/middleware"

	"github.com/gin-gonic/gin"
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

	//g.POST("/login", user.Login)
	//g.POST("/cluster", k8s.ClusterCreate)

	v1 := g.Group("/api/v1/")
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/cpu", handler.CPUCheck)
		v1.GET("/ram", handler.RAMCheck)
		v1.POST("/alert", handler.Alert)
	}

	return g
}
