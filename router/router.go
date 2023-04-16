package router

import (
	"github.com/gin-gonic/gin"
	"go-openai-service/http_handle/hello"
	"net/http"
)

func InitHttpRouter(g *gin.Engine, m ...gin.HandlerFunc) *gin.Engine {

	g.Use()
	g.Use(gin.Recovery())
	g.Use(m...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The router is not exist!")
	})

	g.GET("/hello", hello.Hello)

	return g
}
