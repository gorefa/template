package router

import (
	"net/http"

	"gogin/handler"
	"gogin/handler/user"
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

	v1 := g.Group("/api/v1/")
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/disk", handler.DiskCheck)
		v1.GET("/cpu", handler.CPUCheck)
		v1.GET("/ram", handler.RAMCheck)
	}

	u := g.Group("/api/v1/user")
	{
		u.POST("", user.Create)
		u.GET("", user.List)
	}

	return g
}
